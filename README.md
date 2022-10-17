# Goroutines
Golang concurrency example with the twitter streaming API, using goroutines, channels, waitGroups, and context

Sample output-
pID # 15  - Response error= 429 Too Many Requests
pID # 14  - Response error= 429 Too Many Requests
pID # 11  - Response error= 429 Too Many Requests
pID # 10  - Response error= 429 Too Many Requests
pID # 2  - Response error= 429 Too Many Requests
pID # 1  - Response error= 429 Too Many Requests
pID # 4  - Response error= 429 Too Many Requests
pID # 8  - Response error= 429 Too Many Requests
pID # 20  - Retrieving tweet 1
pID # 5  - Retrieving tweet 1
pID # 16  - Retrieving tweet 1
pID # 3  - Retrieving tweet 1
pID # 9  - Retrieving tweet 1
pID # 7  - Retrieving tweet 1
pID # 6  - Retrieving tweet 1
pID # 19  - Retrieving tweet 1
pID # 17  - Retrieving tweet 1
pID # 13  - Retrieving tweet 1
CURRENT PROGRESS: tweets received = 10 , hashtags = map[]
pID # 12  - Retrieving tweet 1
pID # 18  - Retrieving tweet 1
pID # 13  - Retrieving tweet 2
pID # 9  - Retrieving tweet 2
pID # 20  - Retrieving tweet 2
pID # 17  - Retrieving tweet 2
pID # 9  - Retrieving tweet 3
pID # 9  - Retrieving tweet 4
pID # 12  - Retrieving tweet 2
pID # 13  - Retrieving tweet 3
CURRENT PROGRESS: tweets received = 20 , hashtags = map[#BlastFMRandB:3 #DebateNaBand:1]
pID # 16  - Retrieving tweet 2
pID # 20  - Retrieving tweet 3
pID # 19  - Retrieving tweet 2
pID # 18  - Reader error= EOF
pID # 6  - Retrieving tweet 2
pID # 3  - Retrieving tweet 2
pID # 19  - Retrieving tweet 3
pID # 17  - Retrieving tweet 3
pID # 5  - Retrieving tweet 2
pID # 7  - Retrieving tweet 2
pID # 12  - Retrieving tweet 3
CURRENT PROGRESS: tweets received = 30 , hashtags = map[#BlastFMRandB:6 #DebateNaBand:1]
pID # 20  - Retrieving tweet 4
pID # 17  - Retrieving tweet 4
pID # 19  - Retrieving tweet 4
pID # 5  - Retrieving tweet 3
pID # 7  - Retrieving tweet 3
pID # 7  - Retrieving tweet 4
pID # 6  - Retrieving tweet 3
pID # 19  - Retrieving tweet 5
pID # 3  - Retrieving tweet 3
pID # 16  - Retrieving tweet 3
CURRENT PROGRESS: tweets received = 40 , hashtags = map[#BlastFMRandB:10 #DebateNaBand:4 #LISA:1 #MONEY:1 #SG:1]
pID # 17  - Retrieving tweet 5
pID # 9  - Retrieving tweet 5
pID # 9  - Retrieving tweet 6
pID # 13  - Retrieving tweet 4
pID # 12  - Retrieving tweet 4
pID # 3  - Retrieving tweet 4
pID # 19  - Retrieving tweet 6
pID # 6  - Retrieving tweet 4
pID # 12  - Retrieving tweet 5
pID # 13  - Retrieving tweet 5
CURRENT PROGRESS: tweets received = 50 , hashtags = map[#BlastFMRandB:10 #DebateNaBand:7 #LISA:5 #MONEY:5 #SG:5]
pID # 20  - Retrieving tweet 5
pID # 16  - Retrieving tweet 4
pID # 12  - Retrieving tweet 6
pID # 17  - Retrieving tweet 6
pID # 3  - Retrieving tweet 5
pID # 6  - Retrieving tweet 5
pID # 6  - Retrieving tweet 6
pID # 19  - Retrieving tweet 7
pID # 20  - Retrieving tweet 6
pID # 5  - Retrieving tweet 4
CURRENT PROGRESS: tweets received = 60 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:9 #LISA:6 #MONEY:6 #SG:6]
pID # 7  - Retrieving tweet 5
pID # 3  - Retrieving tweet 6
pID # 9  - Retrieving tweet 7
pID # 3  - Retrieving tweet 7
pID # 3  - Retrieving tweet 8
pID # 6  - Retrieving tweet 7
pID # 17  - Retrieving tweet 7
pID # 6  - Retrieving tweet 8
pID # 12  - Retrieving tweet 7
pID # 13  - Retrieving tweet 6
CURRENT PROGRESS: tweets received = 70 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:9 #LISA:8 #MONEY:8 #SG:8]
pID # 13  - Retrieving tweet 7
pID # 5  - Retrieving tweet 5
pID # 16  - Retrieving tweet 5
pID # 19  - Retrieving tweet 8
pID # 17  - Retrieving tweet 8
pID # 7  - Retrieving tweet 6
pID # 5  - Retrieving tweet 6
pID # 20  - Retrieving tweet 7
pID # 13  - Retrieving tweet 8
pID # 7  - Retrieving tweet 7
CURRENT PROGRESS: tweets received = 80 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:11 #LISA:10 #MONEY:10 #SG:10]
pID # 3  - Retrieving tweet 9
pID # 5  - Retrieving tweet 7
pID # 6  - Retrieving tweet 9
pID # 17  - Retrieving tweet 9
pID # 5  - Retrieving tweet 8
pID # 17  - Retrieving tweet 10
pID # 19  - Retrieving tweet 9
pID # 9  - Retrieving tweet 8
pID # 5  - Retrieving tweet 9
pID # 7  - Retrieving tweet 8
CURRENT PROGRESS: tweets received = 90 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:11 #LISA:10 #MONEY:10 #SG:10 #StopHazaraGenocide:3]
pID # 3  - Retrieving tweet 10
pID # 16  - Retrieving tweet 6
pID # 13  - Retrieving tweet 9
pID # 20  - Retrieving tweet 8
pID # 12  - Retrieving tweet 8
pID # 20  - Retrieving tweet 9
pID # 6  - Retrieving tweet 10
pID # 19  - Retrieving tweet 10
pID # 17  - Retrieving tweet 11
pID # 7  - Retrieving tweet 9
CURRENT PROGRESS: tweets received = 100 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:11 #LISA:11 #MONEY:11 #SG:11 #StopHazaraGenocide:6]
Maximum number of tweets received.  Ending processing.
FINAL PROGRESS: tweets received = 100 , hashtags = map[#BlastFMRandB:11 #DebateNaBand:11 #LISA:11 #MONEY:11 #SG:11 #StopHazaraGenocide:6]
pID # 3  - Context done
pID # 6  - Context done
pID # 9  - Context done
pID # 7  - Context done
pID # 16  - Context done
pID # 5  - Context done
pID # 17  - Context done
pID # 20  - Context done
pID # 13  - Context done
pID # 12  - Context done
pID # 19  - Context done
Processing completed
