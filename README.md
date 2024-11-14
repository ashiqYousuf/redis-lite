# Redis-lite: A Redis-like Database in Go

Redis-lite is a custom Redis database implementation in Go, designed to replicate Redis' core features for **in-memory storage**, **caching**, and **fast data access**. 

## Features

- **In-memory data storage**: Store and retrieve data quickly.
- **Caching**: Efficient caching mechanisms for high-performance systems.
- **Redis-like commands**: Mimics Redis commands for ease of use.

> "Redis-lite is perfect for learning and experimenting with Redis-like data structures and commands."

## Why Use Redis-lite?

Redis-lite is a great tool for developers who want to:
- Learn **database fundamentals** in Go
- Experiment with **Redis-like functionality**
- Understand **in-memory data management**

---


### RESP Protocol (Bulk strings example):

```go

func M() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		log.Fatal("Invalid type, expecting bulk strings only")
	}

	size, _ := reader.ReadByte()
	strSize, err := strconv.ParseInt(string(size), 10, 64)
	if err != nil {
		log.Fatal("invalid type conversion:", err)
	}

	// consume /r/n
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	// consume /r/n
	reader.ReadByte()
	reader.ReadByte()

	fmt.Println(string(name))
}
```