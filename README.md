## hashpipe / hpipe

Pipe CSV into it, with the fields you want obscured, and they'll be hashed

```
cat somecsv.csv | hpipe -fields 0,2,4 > hashed.csv
```

WIP