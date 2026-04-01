## transmute

A command-line unit converter. Supports distance, weight, volume, temperature, velocity, area, pressure, energy, data, and time — with metric, imperial, historical, and astronomical units.

```sh
go install github.com/MawCeron/transmute@latest
```

```sh
transmute 10 km mi
transmute -c weight 70 kg lb
transmute -c temperature 100 c f
transmute -c data 1 gb mb
transmute -c distance --list
```

Full unit reference: [docs/](docs/)
