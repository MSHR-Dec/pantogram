# Create ERD
```
docker run \
    -v "$PWD/doc:/output" \
    --net="pantogram" \
    schemaspy/schemaspy:snapshot \
        -t mysql \
        -host mysql:3306 \
        -db pantogram \
        -u root \
        -p root \
        -connprops useSSL\\=false \
        -s pantogram
```