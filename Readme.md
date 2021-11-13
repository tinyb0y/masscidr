# MassCIDR

Takes multiple CIDR as input and converts them into a masscan script file.

### Usage Example

```bash
$ ./main
Usage: masscidr -i cidr.txt -o outfile
```

Input file
```bash
$ cat cidr.txt

192.168.56.0/24
192.168.57.0/24
```
Output file
```bash
$ cat out.txt 
masscan -p1-65535 --rate=10000 192.168.56.0/24
masscan -p1-65535 --rate=10000 192.168.57.0/24
```
