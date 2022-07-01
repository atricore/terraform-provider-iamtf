# Logging / Debugging the provider

JOSSO_API_TRACE=true for network traffic dump

For terraform logging, valid values: [TRACE DEBUG INFO WARN ERROR OFF]

TF_LOG
TF_LOG_PROVIDER
TF_LOG_PATH
TF_LOG_ACC_PATH

## Enable JOSSO api traffic log

This requiers TF_LOG_PROVIDER=TRACE

`export JOSSO_API_TRACE=true`

## Enable Terraform log

### Option 1

`export TF_LOG=TRACE TF_LOG_PATH=/tmp/tf.log TF_LOG_ACC_PATH=/tmp/tf_acc.log TF_LOG_PROVIDER=TRACE`

### Option 2

`export TF_LOG=TRACE TF_LOG_PROVIDER=TRACE`


## Visual Studio Code

### Environment

Create .env file in project root dir.

```
TF_ACC=1
TFE_PARALLELISM=1
```

Creaate **launch.json**


```json
{ 
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch a test",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "./josso",
            "args": [
                "-test.v",
                "-test.run",
                "^${selectedText}$"
            ],
            "env": {
                "TF_ACC": "1"
            },
            "buildFlags": "-v",
            "showLog": true,
            "envFile": "${workspaceFolder}/.env"
        }

    ]
}
```

TODO: Add more docs
