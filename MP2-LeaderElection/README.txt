MP2 Owen Markley

Running Code:
  - Code should run using command 'go run raft_election.go'
  - Program is currently configured for 5 nodes
  - Simulating Failure:
    - In order to simulate a failure, set the variable hbBeforeFail to an integer, after that 
      number of rounds of sending heartbeats, the currently elected leader will then fail
      making room for another leader to be chosen.

Output:
  - Output from the program should be printed to the terminal
  - Sample outputs are listed at the bottom of the file
  - A message is generated each time a node receives a signal:
    - Signals could be either a clock signal, indicating that an election timer
      has run out, or a message sent from another node (vote request, vote, 
      heartbeat, heartbeat acknowladgement)
    - These will generally show what the message is, and from who it was received
    - NOTE: Once the election is complete, the messages become flooded with heartbeats
      and their respective ack's relatively quickly

Sample Output:

Electing a Leader:
Hello, World!
ConsensusModule - ID: 0, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 1, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 2, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 3, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 4, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
3 started election after timeout at 181ms
4 received vote request from 3
3 received a vote, for a total of  2
2 received vote request from 3
3 received a vote, for a total of  3
3 is now Leader
0 received vote request from 3
1 received vote request from 3
4 received heartbeat from 3
3 received ack from 4
2 received heartbeat from 3
3 received ack from 2
1 received heartbeat from 3
3 received ack from 1
0 received heartbeat from 3
3 received ack from 0
4 received heartbeat from 3
0 received heartbeat from 3
3 received ack from 4
3 received ack from 0
1 received heartbeat from 3

*Running was cut off after leader was chosen and only heartbeats are exchanged*

Simulating Failure:
Hello, World!
ConsensusModule - ID: 0, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 1, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 2, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 3, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
ConsensusModule - ID: 4, Peers: [0 1 2 3 4], Term: 0, Vote: -1, Timeout: 0001-01-01 00:00:00 +0000 UTC
3 started election after timeout at 181ms
4 received vote request from 3
3 received a vote, for a total of  2
0 received vote request from 3
3 received a vote, for a total of  3
3 is now Leader
1 received vote request from 3
2 received vote request from 3
4 received heartbeat from 3
3 received ack from 4
0 received heartbeat from 3
1 received heartbeat from 3
2 received heartbeat from 3
3 received ack from 0
3 received ack from 1
3 received ack from 2
4 received heartbeat from 3
1 received heartbeat from 3
0 received heartbeat from 3
3 received ack from 4
3 received ack from 1
2 received heartbeat from 3
3 received ack from 0
3 received ack from 2
4 received heartbeat from 3
3 received ack from 4
0 received heartbeat from 3
3 received ack from 0
2 received heartbeat from 3
3 received ack from 2
1 received heartbeat from 3
3 received ack from 1
0 started election after timeout at 197ms
4 received vote request from 0
0 received a vote, for a total of  2
1 received vote request from 0
0 received a vote, for a total of  3
0 is now Leader
2 received vote request from 0
4 received heartbeat from 0
0 received ack from 4
1 received heartbeat from 0
0 received ack from 1
3 received ack from 0
2 received heartbeat from 0
0 received ack from 2
4 received heartbeat from 0
0 received ack from 4


