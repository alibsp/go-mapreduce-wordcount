# Go MapReduce Word Count

A hands-on exercise in Go to understand the MapReduce programming model by implementing a word-count algorithm.

The exercise starts with a sequential implementation and explores concurrent processing using goroutines and channels.

## Overview

MapReduce is a programming model for processing data through three main stages:

- **Map:** Splits the input text into words and emits key-value pairs.
- **Shuffle:** Groups values associated with the same key.
- **Reduce:** Aggregates the grouped values to calculate the final word counts.

This exercise implements these stages in Go and compares sequential and concurrent execution.

## Architecture

The processing pipeline consists of three stages:

```text
          Input Text
               |
               v
             Map
               |
               v
            Shuffle
               |
               v
             Reduce
               |
               v
         Word Frequencies
```

### Map

Splits the input text into words and emits a `KeyValue` pair for each word.

Example:

```text
Input:  "Go is great Go"
Output: (go, 1), (is, 1), (great, 1), (go, 1)
```

### Shuffle

Groups values by their keys.

Example:

```text
(go, 1), (is, 1), (great, 1), (go, 1)
```

becomes:

```text
go    -> [1, 1]
is    -> [1]
great -> [1]
```

### Reduce

Aggregates the values for each key to produce the final word counts.

Example:

```text
go    -> 2
is    -> 1
great -> 1
```

## Implementations

This exercise explores the following approaches:

### 1. Sequential MapReduce

Executes Map, Shuffle, and Reduce sequentially in a single goroutine.

### 2. Concurrent MapReduce — Method 1

- Splits the input text into partitions.
- Runs multiple Mapper workers using goroutines.
- Transfers worker results through a channel.
- Collects all mapped key-value pairs in an intermediate slice.
- Executes Shuffle and Reduce after mapping is complete.

### 3. Concurrent MapReduce — Method 2

- Splits the input text into partitions.
- Runs multiple Mapper workers using goroutines.
- Transfers worker results through a channel.
- Groups each worker's results immediately after receiving them.
- Executes Reduce after all mapping results have been processed.

Method 2 avoids storing all raw mapped pairs in an additional intermediate slice.

**Note:** Both concurrent methods currently operate within a single process and on a single machine. They are concurrent implementations, not distributed MapReduce systems.

## Requirements

- Go

## Getting Started

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/go-mapreduce-wordcount.git
cd go-mapreduce-wordcount
```

Run the program:

```bash
go run .
```

Replace `YOUR_USERNAME` with your GitHub username.

If the program processes an input file, make sure the file exists at the path expected by the application.

## Performance Comparison

The exercise includes a basic timing comparison between sequential and concurrent execution.

The comparison checks that both implementations produce identical results and measures their execution times.

Performance depends on several factors, including:

- Input size
- Number of Mapper workers
- Goroutine scheduling overhead
- Memory allocation and garbage collection
- CPU resources

For small inputs, concurrent execution may be slower than sequential execution because of its additional overhead.

A more reliable performance evaluation should use repeated measurements or Go benchmarks.

## Current Limitations

- The entire input file is read into memory.
- Input partitioning is performed before Mapper workers start.
- Shuffle and Reduce run in the main goroutine.
- Worker communication uses an in-process channel.
- Word normalization currently converts words to lowercase but does not remove punctuation.
- The current implementation does not distribute work across multiple machines.

## Future Improvements

Possible extensions include:

- Reading large files in chunks.
- Handling word boundaries when partitioning files.
- Improving memory efficiency.
- Benchmarking different numbers of Mapper workers.
- Implementing concurrent Reduce workers.
- Partitioning keys across multiple Reducers.
- Exploring inter-process communication and network-based workers.
- Extending the implementation toward distributed MapReduce.

## Learning Goals

The main goals of this exercise are to:

- Understand the MapReduce programming model.
- Practice working with slices and maps in Go.
- Learn how to use goroutines and channels.
- Compare sequential and concurrent processing.
- Explore performance and memory trade-offs.
- Build a foundation for studying distributed data processing.

## License

This project is available under the MIT License.

If you choose the MIT License, add a `LICENSE` file containing the appropriate license text.
