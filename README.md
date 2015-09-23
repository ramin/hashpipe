## hashpipe / hpipe

Pipe CSV into it, with the fields you want obscured, and they'll be hashed

```
cat test.csv | hpipe -fields 0,2,4 > hashed.csv

# Where the input of test.csv might be
# sensistive,ok,not bad,please obscure this,whatevr,let it live
# Thus the output in hashed.csv would then be
# 32ff2a55edb63ab48c9a97a87009a7b6,ok,not bad,35b907b615ce38f48dedfd2d057e008d,whatevr,let it live

```