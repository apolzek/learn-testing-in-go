
```bash
echo "POST http://localhost:8080/multiply
# sending with payload
@data.json" | vegeta attack -duration=10s | tee results.bin | vegeta report
```

```bash
echo "POST http://localhost:8080/multiply
@data.json" | vegeta attack -rate=100 -duration=30s > attack-results.bin

cat attack-results.bin | vegeta plot > plot.html
```


