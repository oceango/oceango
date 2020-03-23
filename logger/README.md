# How to use

```
import "github.com/oceango/logger"

log := logger.GetLogger()
log.Output("example log")   // output to Console
log.Write("example log")    // output to file
log.ALL("example log")    // output to file and console
log.FULL(logger.GET, "example log", false)  // args true write to file 
```

# Log file store where?

> {YourPath} / logs