# ClassMITM-MQTT
This repository contains code to simulate a Man-in-the-Middle (MITM) attack using MQTT for a class at Inteli University. The simulation includes a producer, a listener, and a Man-in-the-Middle component. The messages exchanged between the producer and the listener are encrypted using the Caesar cipher.

# Video demo
https://github.com/user-attachments/assets/7865fa88-1a95-4ea7-ba13-0b66cabe5631

## How to Run

> [!IMPORTANT]
> We strongally recommand the use os `just`. If you don't have it, please check: [just installation](https://github.com/casey/just)

## Prerequisites

1. **Environment Variables:**
   - All executables depend on a `.env` file for configuration.
   - Refer to the provided `.env.example` to see which values are required.
   - You can use the `-env` flag to specify the path to your `.env` file if it's not in the default location (`./.env`).

2. **Required Flags:**
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

### Commands

| Command                         | Description                                     |
|---------------------------------|-------------------------------------------------|
| `just compile`                  | Compiles all executables into `./build/bin`.    |
| `just run <executable> [flags]` | Runs the specified executable with optional flags. Available: `mitm`, `listener`, or `publisher` |
