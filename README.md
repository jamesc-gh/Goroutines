# Goroutines
Golang concurrency example with the Twitter streaming API, using goroutines, channels, waitGroups, and context

Sample output-
```
go run .
pID # 7  - Response error = 429 Too Many Requests
pID # 12  - Response error = 429 Too Many Requests
pID # 8  - Retrieving tweet 1
pID # 1  - Retrieving tweet 1
pID # 1  - Pause threshold surpassed.  Pushing skipPause message.
pID # 13  - Retrieving tweet 1
pID # 13  - SkipPause message read from the queue.
pID # 13  - Retrieving tweet 2
pID # 20  - Retrieving tweet 1
pID # 16  - Retrieving tweet 1
pID # 3  - Retrieving tweet 1
pID # 3  - Pause threshold surpassed.  Pushing skipPause message.      
pID # 5  - Retrieving tweet 1
pID # 5  - SkipPause message read from the queue.
pID # 6  - Retrieving tweet 1
pID # 18  - Retrieving tweet 1
CURRENT PROGRESS: tweets received = 10 , hashtags = map[]
pID # 5  - Retrieving tweet 2
pID # 15  - Retrieving tweet 1
pID # 19  - Retrieving tweet 1
pID # 2  - Retrieving tweet 1
pID # 9  - Retrieving tweet 1
pID # 10  - Retrieving tweet 1
pID # 4  - Retrieving tweet 1
pID # 15  - Retrieving tweet 2
pID # 14  - Retrieving tweet 1
pID # 11  - Retrieving tweet 1
pID # 11  - Pause threshold surpassed.  Pushing skipPause message.     
pID # 17  - Retrieving tweet 1
CURRENT PROGRESS: tweets received = 20 , hashtags = map[]
pID # 17  - SkipPause message read from the queue.
pID # 17  - Retrieving tweet 2
pID # 14  - Retrieving tweet 2
pID # 14  - Pause threshold surpassed.  Pushing skipPause message.
pID # 10  - Retrieving tweet 2
pID # 10  - SkipPause message read from the queue.
pID # 10  - Retrieving tweet 3
pID # 16  - Retrieving tweet 2
pID # 5  - Retrieving tweet 3
pID # 8  - Retrieving tweet 2
pID # 5  - Retrieving tweet 4
pID # 4  - Reader error = EOF
pID # 15  - Retrieving tweet 3
CURRENT PROGRESS: tweets received = 30 , hashtags = map[#AMas:1]
pID # 17  - Retrieving tweet 3
pID # 9  - Retrieving tweet 2
pID # 19  - Retrieving tweet 2
pID # 16  - Retrieving tweet 3
pID # 13  - Retrieving tweet 3
pID # 18  - Reader error = EOF
pID # 6  - Retrieving tweet 2
pID # 2  - Reader error = EOF
pID # 17  - Retrieving tweet 4
pID # 20  - Retrieving tweet 2
pID # 20  - Pause threshold surpassed.  Pushing skipPause message.
pID # 1  - Retrieving tweet 2
pID # 1  - SkipPause message read from the queue.
pID # 1  - Retrieving tweet 3
CURRENT PROGRESS: tweets received = 40 , hashtags = map[#AMas:1]       
pID # 1  - Retrieving tweet 4
pID # 8  - Retrieving tweet 3
pID # 15  - Retrieving tweet 4
pID # 3  - Retrieving tweet 2
pID # 17  - Retrieving tweet 5
pID # 9  - Retrieving tweet 3
pID # 5  - Retrieving tweet 5
pID # 10  - Retrieving tweet 4
pID # 11  - Retrieving tweet 2
CURRENT PROGRESS: tweets received = 50 , hashtags = map[#AMas:4]
pID # 9  - Retrieving tweet 4
pID # 14  - Retrieving tweet 3
pID # 11  - Retrieving tweet 3
pID # 3  - Retrieving tweet 3
pID # 15  - Retrieving tweet 5
pID # 15  - Pause threshold surpassed.  Pushing skipPause message.
pID # 6  - Retrieving tweet 3
pID # 6  - SkipPause message read from the queue.
pID # 6  - Retrieving tweet 4
pID # 10  - Retrieving tweet 5
pID # 16  - Retrieving tweet 4
pID # 8  - Retrieving tweet 4
CURRENT PROGRESS: tweets received = 60 , hashtags = map[#AMas:4]
pID # 9  - Retrieving tweet 5
pID # 3  - Retrieving tweet 4
pID # 6  - Retrieving tweet 5
pID # 17  - Retrieving tweet 6
pID # 3  - Retrieving tweet 5
pID # 13  - Retrieving tweet 4
pID # 11  - Retrieving tweet 4
pID # 19  - Retrieving tweet 4
pID # 14  - Retrieving tweet 4
pID # 19  - Retrieving tweet 5
CURRENT PROGRESS: tweets received = 70 , hashtags = map[#AMas:7 #BallonDor:1 #FiersdetreBleus:1]
pID # 19  - Retrieving tweet 6
pID # 20  - Retrieving tweet 3
pID # 1  - Retrieving tweet 5
pID # 9  - Retrieving tweet 6
pID # 5  - Retrieving tweet 6
pID # 5  - Pause threshold surpassed.  Pushing skipPause message.
pID # 8  - Retrieving tweet 5
pID # 8  - SkipPause message read from the queue.
pID # 8  - Retrieving tweet 6
pID # 20  - Retrieving tweet 4
pID # 6  - Retrieving tweet 6
pID # 8  - Retrieving tweet 7
CURRENT PROGRESS: tweets received = 80 , hashtags = map[#AMas:8 #BallonDor:1 #FiersdetreBleus:1]
pID # 16  - Retrieving tweet 5
pID # 1  - Retrieving tweet 6
pID # 10  - Retrieving tweet 6
pID # 6  - Retrieving tweet 7
pID # 20  - Retrieving tweet 5
pID # 15  - Retrieving tweet 6
pID # 10  - Retrieving tweet 7
pID # 13  - Retrieving tweet 5
pID # 19  - Retrieving tweet 7
pID # 11  - Retrieving tweet 5
CURRENT PROGRESS: tweets received = 90 , hashtags = map[#AMas:11 #BallonDor:1 #FiersdetreBleus:1]
pID # 19  - Retrieving tweet 8
pID # 17  - Retrieving tweet 7
pID # 19  - Retrieving tweet 9
pID # 3  - Retrieving tweet 6
pID # 6  - Retrieving tweet 8
pID # 6  - Pause threshold surpassed.  Pushing skipPause message.
pID # 9  - Retrieving tweet 7
pID # 9  - SkipPause message read from the queue.
pID # 9  - Retrieving tweet 8
pID # 1  - Retrieving tweet 7
pID # 16  - Retrieving tweet 6
pID # 14  - Retrieving tweet 5
CURRENT PROGRESS: tweets received = 100 , hashtags = map[#AMas:11 #BallonDor:1 #FiersdetreBleus:1 #devtober:1 #pixelart:1]
Maximum number of tweets received.  Ending processing.
FINAL PROGRESS: tweets received = 100 , hashtags = map[#AMas:11 #BallonDor:1 #FiersdetreBleus:1 #devtober:1 #pixelart:1]
pID # 17  - Context done
pID # 13  - Context done
pID # 8  - Context done
pID # 10  - Context done
pID # 11  - Context done
pID # 15  - Context done
pID # 20  - Context done
pID # 5  - Context done
pID # 3  - Context done
pID # 19  - Context done
pID # 1  - Context done
pID # 16  - Context done
pID # 9  - Context done
pID # 14  - Context done
pID # 6  - Context done
Processing completed
```
