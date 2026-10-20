# Concurrent MapReduce Word Count in Go

A small educational implementation of the **MapReduce** programming model in Go. It counts words across multiple text files and uses goroutines and channels to coordinate the Map, Shuffle/Group, and Reduce stages.

## Features

- Processes multiple input files using concurrent Mapper goroutines.
- Normalizes words to lowercase, so `Go` and `go` are counted together.
- Splits text on whitespace and punctuation, including hyphens (`-`) and repeated hyphens (`--`).
- Groups intermediate values by word during Shuffle/Group.
- Assigns each word to one of `nReduce` partitions using `hash(key) % nReduce`.
- Runs one Reduce worker goroutine per partition.
- Collects final counts through a channel.
- Uses Go's standard library; no third-party dependencies are required for hashing or word tokenization.

## How It Works

1. **Map:** Each `MapWorker` reads one input file and emits a `KeyValue` pair for every word, such as `{"hello", 1}`.
2. **Shuffle/Group:** The intermediate pairs from all Mapper workers are grouped by word. For example, `hello → [1, 1, 1]`.
3. **Partition:** Each word is assigned to a partition using `hash(key) % nReduce`. A given word always maps to the same partition for the chosen hash function and partition count.
4. **Reduce:** One `ReduceWorker` goroutine is started for each partition. It sums the values associated with each word and sends a final `KeyValue` result through a channel.
5. **Collect:** The main workflow receives the final results and stores them in a `map[string]int`.

## Requirements

- Go installed on your machine.
- One or more readable text files to use as input.

The implementation uses Go's standard library, including `hash/fnv`, `os`, `strings`, and `unicode`.

## Usage

Set the input file paths and the number of Reduce partitions in your `main` function, then call `MapReduce_Concurrent`. For example:

```go
func main() {
    files := []string{
        "input1.txt",
        "input2.txt",
        "input3.txt",
    }

    nReduce := 3
    MapReduce_Concurrent(files, nReduce)
}
```

The filenames above are examples; replace them with paths to files that exist in your project. If your project already has a `main` function, update it rather than adding a second one.

From the project root, run:

```bash
go run .
```

This command assumes the project is configured as a Go module. `nReduce` must be greater than zero.

## Example

Suppose the input files collectively contain:

```text
Hello, world!
hello Go
GO-world
```

After lowercasing and splitting on punctuation, the expected counts are:

```text
go 2
hello 2
world 2
```

The order of printed words is not guaranteed in the current implementation because Go map iteration order is unspecified.

## Main Functions and Types

| Function / Type | Responsibility |
|---|---|
| `KeyValue` | Represents a word and its intermediate or final count. |
| `Map` | Tokenizes text and emits one key-value pair per word. |
| `MapWorker` | Reads an input file and runs the Map stage. |
| `ShuffleInto` | Groups intermediate values by word. |
| `Partition` | Maps a word to a partition using a hash function. |
| `Reduce` | Sums the values for one word. |
| `ReduceWorker` | Reduces all words in one partition and sends results through a channel. |
| `RunConcurrent_method2` | Coordinates the Map, Shuffle/Group, Partition, and Reduce stages. |
| `MapReduce_Concurrent` | Runs the workflow and prints the resulting counts. |

## Current Limitations / Possible Improvements

- If an input file cannot be read, `MapWorker` prints an error and sends an empty result. A future version could propagate errors to the caller instead.
- The current implementation prints results to the terminal. Writing alphabetically sorted results to `output.txt` is a possible next step and is not implemented in the current code shown here.
- To print results in alphabetical order, collect the keys from the result map and sort them with `sort.Strings` before printing.
- All punctuation is treated as a delimiter. For example, `don't` is tokenized as `don` and `t`.

## Learning Goals

This project is intended to practice:

- Goroutines and concurrent task execution.
- Channels for communication between goroutines.
- The MapReduce workflow: Map, Shuffle/Group, and Reduce.
- Hash-based partitioning.
- Collecting results from concurrent workers.
