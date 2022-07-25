# Profiling

To profile a program, you can use the following command:

```bash
make build-profile-cpu
```

```bash
make build-profile-mem
```

```bash
make build-race
```

To see the information, you can use the `pprof` tool.

```bash
go tool pprof <mem.pprof or cpu.pprof>
```

To see the information on a web page, you can use the `-web` flag.

```bash
go tool pprof -web <mem.pprof or cpu.pprof>
```

<https://go.dev/blog/pprof>

<https://youtu.be/nok0aYiGiYA>

<https://youtu.be/2h_NFBFrciI>
