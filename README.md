# configReader

The utility helps to read json configuration file that may contain your configs for your appliction. Basically it is a common workflow to process configuration files.

# Setup
- put the following commands to cmd to create binary.
```sh
    go build .
```

# Usage
```sh
    .\configreader.exe
        Try to unmarshal configs...
        We recieved this structure:
{pass_linux pass_windows true machine1 joshua 8 machine3 joseph 9}
```

# Acces to docs
- puth to cmd and press enter.


# Errors
* code 1 - reading configs problem.
* code 2 - unmarshal problem.