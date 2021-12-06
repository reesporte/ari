# ari
`ari` will lex, parse and evaluate arithmetic expressions from a file
```bash
$ cat example.input
420 + 69 + 69 * 38 + 400 - 69.32 / 800.4343434
400 / 54 + 38.2 - 77 
$ ari example.input
420 + 69 + 69 * 38 + 400 - 69.32 / 800.4343434 = 3510.9133970192916
400 / 54 + 38.2 - 77  = -31.39259259259259
```

