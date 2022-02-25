// Raft Election Implementation
// Owen Markley CS 499 - Distributed Systems
package main

import (
  "strconv"
	"fmt"
	"time"
	"math/rand"
	"sync"
)

// For simplicity, log entries are comprised of only ints
type LogEntry struct {
	Term int
}

type ConsensusModule struct {
	// Mutex is needed in order for safe concurrent access to a ConsensusModule
	mu sync.Mutex
	// id is the server ID of this ConsensusModule
	id          int
	// peerIds contains the IDs of peers in the cluster
	peerIds     *[]int
  //IMPORTANT: peerIds needs to be replaced with a list of channels rather than a list of ids
  chans       *[5]chan string
  // State indicates follower:1, candidate:2, or leader:3
  state       int
	// Components as defined in Raft paper:
	// Persistent State
	currentTerm int
	votedFor    int
  
  electionResetEvent time.Time
}

func (cm ConsensusModule) String() string {
  return fmt.Sprintf("ConsensusModule - ID: %v, Peers: %v, Term: %v, Vote: %v, Timeout: %v", cm.id, *cm.peerIds, cm.currentTerm, cm.votedFor, cm.electionResetEvent)
}

func (cm *ConsensusModule) electionTimeout() time.Duration {
	return time.Duration(150+rand.Intn(150)) * time.Millisecond
}

func (cm *ConsensusModule) runElectionTimer() {
  self := cm.id
	timeoutDuration :=  cm.electionTimeout()
  timer := time.NewTimer(timeoutDuration)
  heartbeatTimer := time.NewTimer(time.Millisecond * 100)
  voteCount := 0
  hbBeforeCrash := 4
  for {
    select{
    // Election timeout
    case <-timer.C:
      // Election timeout in follower state
      if cm.state == 1 {
        fmt.Println(self, "started election after timeout at", timeoutDuration)
        cm.state = 2
        voteCount = voteCount + 1
        go cm.sendVoteRequest()
        timer.Reset(timeoutDuration)
      } else if cm.state == 2 { // Election timeout in candidate state
        voteCount = 0
        fmt.Println(self, "restarted election after timeout at", timeoutDuration)
        timer.Reset(timeoutDuration)
      } else if cm.state == 3 { // Election timeout in leader (do nothing)
        timer.Reset(timeoutDuration)
      }
      
    // Message received
    case input := <-(*cm.chans)[cm.id]:
      // Message is vote request
      if string(input[0:2]) == "vr" {
        // Only vote if in follower mode
        if cm.state == 1 && cm.votedFor == -1 {
          if target, err := strconv.Atoi(string(input[2])); err == nil {
            fmt.Println(self, "received vote request from", target)
            go cm.vote(target)
          }
          timer.Reset(timeoutDuration)
        }
      }
      // Message is vote
      if string(input[0:2]) == "vf" && cm.state == 2 {
        voteCount = voteCount + 1
        fmt.Println(self, "received a vote, for a total of ", voteCount)
        // Candidate has received enough votes to win
        if voteCount > 2 {
          fmt.Println(self, "is now Leader")
          cm.state = 3
          voteCount = 0
        }
      }
      // Message is heartbeat
      if string(input[0:2]) == "hb" {
        //return
        if target, err := strconv.Atoi(string(input[2])); err == nil {
          if cm.state == 1 {
            fmt.Println(self, "received heartbeat from", target)
            cm.votedFor = -1
            go cm.sendHeartbeatAck(target)
            timer.Reset(timeoutDuration)
          } else if cm.state == 2 {
            fmt.Println(self, "received heartbeat from", target)
            cm.votedFor = -1
            voteCount = 0
            cm.state = 1
            go cm.sendHeartbeatAck(target)
            timer.Reset(timeoutDuration)
          } else if cm.state == 3 {
            fmt.Println(self, "received ack from", target)
          }
        }
      }
    case <-heartbeatTimer.C:
      if cm.state == 3 {
        hbBeforeCrash = hbBeforeCrash - 1
        if hbBeforeCrash < 1 {
          //cm.state = 1
        } else {
          go cm.sendAllHeartbeats()
        }
      }
      heartbeatTimer.Reset(time.Millisecond * 100)
    }
    
  }
}

func (cm ConsensusModule) sendVoteRequest(){
  self := cm.id
  for i := range (*cm.chans) {
    if i != self {
      (*cm.chans)[i] <- fmt.Sprintf("vr%v", self)
    }
  }
}

func (cm *ConsensusModule) vote(request int){
  cm.votedFor = request
  (*cm.chans)[request] <- fmt.Sprintf("vf%v", request)
}

func (cm ConsensusModule) sendHeartbeat(target int) {
  message := fmt.Sprintf("hb%v", cm.id)
  selectedNode := (*cm.chans)[target]
  selectedNode <- message
}

func (cm ConsensusModule) sendHeartbeatAck(target int) {
  (*cm.chans)[target] <- fmt.Sprintf("hb%v", cm.id)
}

func (cm ConsensusModule) sendAllHeartbeats() {
  time.Sleep(time.Millisecond * 100)
  self := cm.id
  for i := range *cm.chans {
    if i != self {
      go cm.sendHeartbeat(i)
    }
  }
}

func main() {
	fmt.Println("Hello, World!")
  
  // Initializing nodes (creating channels, declaring and defining nodes)...
  nodeIds := []int{0, 1, 2, 3, 4}
  nodes := make([]ConsensusModule, 5)
  var chans [5]chan string
  for i := range chans {
    chans[i] = make(chan string)
  }
  
  for i := range nodes {
    node := ConsensusModule{id: i, peerIds: &nodeIds, chans: &chans, state: 1, votedFor: -1}
    nodes[i] = node
    fmt.Println(node)
  }
  
  go nodes[0].runElectionTimer()
  go nodes[1].runElectionTimer()
  go nodes[2].runElectionTimer()
  go nodes[3].runElectionTimer()
  go nodes[4].runElectionTimer()
  
  var input string
  fmt.Scanln(&input)
}
