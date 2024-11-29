# ClassMITM-MQTT
This repository contains code to simulate a Man-in-the-Middle (MITM) attack using MQTT for a class at Inteli University. The simulation includes a producer, a listener, and a Man-in-the-Middle component. The messages exchanged between the producer and the listener are encrypted using the Caesar cipher.

# Demo video
## Availability
https://github.com/user-attachments/assets/8e23d8e7-5bb5-4bc5-a3f7-78ca7dd84d7b

## Confidentiality and Integrity
https://github.com/user-attachments/assets/7865fa88-1a95-4ea7-ba13-0b66cabe5631

## How to Run

> [!IMPORTANT]
> We strongly recommand the use `just`. If you don't have it, please check: [just installation](https://github.com/casey/just)

## Prerequisites

1. **Required apps**
   - [golang](https://go.dev/)
2. **Environment Variables:**
   - All executables depend on a `.env` file for configuration.
   - Refer to the provided `.env.example` to see which values are required.
   - You can use the `-env` flag to specify the path to your `.env` file if it's not in the default location (`./.env`).

3. **Required Flags:**
   - The `listener` executable **requires** the `-shift` flag.

## With Justfile

The `justfile` provides the following recipes to streamline the workflow:

#### 1. Compile the Executables

Before running any of the executables, you need to compile them:

```bash
just compile
```

This will:
- Create the necessary `./build/bin` directory if it doesn’t already exist.
- Compile the `mitm`, `listener`, and `publisher` executables into `./build/bin`.

#### 2. Run an Executable

To run an executable, use the `run` recipe:

```bash
just run <executable> [flags...]
```

Replace `<executable>` with one of the available options: `mitm`, `listener`, or `publisher`.

#### 3. Examples of Running Each Executable

1. **Run `mitm`:**

Obs.: The man-in-the-middle accepts the `-hijack` flag to intercept the message without forwarding it to the listener.

```bash
just run mitm
```

This runs the `mitm` executable with the specified `.env` file.

2. **Run `listener` (with `-shift` flag):**

```bash
just run listener -shift 9
```

The `-shift` flag is **required** for `listener`. Replace `9` with the desired shift value for the caeser cipher.

3. **Run `publisher`:**

```bash
just run publisher
```

#### Commands

| Command                         | Description                                     |
|---------------------------------|-------------------------------------------------|
| `just compile`                  | Compiles all executables into `./build/bin`.    |
| `just run <executable> [flags]` | Runs the specified executable with optional flags. Available: `mitm`, `listener`, or `publisher` |

## Without Justfile

If you prefer not to use the `Justfile`, you can manually compile and run the executables by following these steps:

### 1. Compiling the Executables

Before running any executable, you need to compile the binaries. Use the following commands to compile each executable.

1. Compile each executable:

- **`mitm`**:
  ```bash
  go build -o ./mitm ./cmd/mitm/main.go
  ```

- **`listener`**:
  ```bash
  go build -o ./listener ./cmd/listener/main.go
  ```

- **`publisher`**:
  ```bash
  go build -o ./publisher ./cmd/publisher/main.go
  ```

### 2. Running the Executables

After compiling, you can execute the binaries. Refer to the following examples:

#### a) Running `mitm`

The `mitm` executable can be run directly. To intercept messages without forwarding them to the listener, use the `-hijack` flag.

```bash
./mitm
```

#### b) Running `listener`

The `listener` requires the `-shift` flag to specify the Caesar cipher shift value.

```bash
./listener -shift 9
```

Replace `9` with the desired shift value.

#### c) Running `publisher`

To run the `publisher`, execute the binary directly:

```bash
./publisher
```

---

### Summary of Commands

| Command                                | Description                                     |
|----------------------------------------|-------------------------------------------------|
| `go build -o ./<target> ./cmd/<target>/main.go`| Compiles the specified executable.             |
| `./mitm`                               | Runs the `mitm` executable. Accept `-hijack` flag.  |
| `./listener -shift <value>`            | Runs the `listener` with the specified shift value. |
| `./publisher`                          | Runs the `publisher` executable.               |
