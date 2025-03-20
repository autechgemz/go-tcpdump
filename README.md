# Gotcap 

This program captures packets on a specified network interface and displays TCP packet information.

## Prerequisites

- `libpcap` library must be installed on your system:

## Usage

1. Install required dependencies:
   ```
   go get github.com/google/gopacket
   go get github.com/google/gopacket/pcap
   ```

2. Run the program:
   ```
   go run main.go -i <interface> -f <BPF filter>
   ```

   - `-i`: Network interface to capture packets from (required)
   - `-f`: BPF filter (required)
   - `-d`: Enable payload debug output (optional)

## Example Output

```
Time: 2025-03-20T21:17:58.917036+09:00, IPv4: 172.17.0.2      -> 192.168.56.201 , TCP: 35716 ->    25, Size:       74 B, Seq: 3632611744, Flags: [SYN --- --- --- --- ---]
Time: 2025-03-20T21:17:58.918845+09:00, IPv4: 192.168.56.201  -> 172.17.0.2     , TCP:    25 -> 35716, Size:       74 B, Seq: 1436965257, Flags: [SYN ACK --- --- --- ---]
Time: 2025-03-20T21:17:58.918885+09:00, IPv4: 172.17.0.2      -> 192.168.56.201 , TCP: 35716 ->    25, Size:       66 B, Seq: 3632611745, Flags: [--- ACK --- --- --- ---]
Time: 2025-03-20T21:17:58.943353+09:00, IPv4: 192.168.56.201  -> 172.17.0.2     , TCP:    25 -> 35716, Size:      103 B, Seq: 1436965258, Flags: [--- ACK PSH --- --- ---]
Time: 2025-03-20T21:17:58.943381+09:00, IPv4: 172.17.0.2      -> 192.168.56.201 , TCP: 35716 ->    25, Size:       66 B, Seq: 3632611745, Flags: [--- ACK --- --- --- ---]
Time: 2025-03-20T21:22:58.604481+09:00, IPv4: 192.168.56.201  -> 172.17.0.2     , TCP:    25 -> 35716, Size:      119 B, Seq: 1436965295, Flags: [--- ACK PSH --- --- ---]
Time: 2025-03-20T21:22:58.604511+09:00, IPv4: 172.17.0.2      -> 192.168.56.201 , TCP: 35716 ->    25, Size:       66 B, Seq: 3632611745, Flags: [--- ACK --- --- --- ---]
Time: 2025-03-20T21:22:58.604640+09:00, IPv4: 192.168.56.201  -> 172.17.0.2     , TCP:    25 -> 35716, Size:       66 B, Seq: 1436965348, Flags: [--- ACK --- FIN --- ---]
Time: 2025-03-20T21:22:58.604749+09:00, IPv4: 172.17.0.2      -> 192.168.56.201 , TCP: 35716 ->    25, Size:       66 B, Seq: 3632611745, Flags: [--- ACK --- FIN --- ---]
Time: 2025-03-20T21:22:58.605232+09:00, IPv4: 192.168.56.201  -> 172.17.0.2     , TCP:    25 -> 35716, Size:       66 B, Seq: 1436965349, Flags: [--- ACK --- --- --- ---]
```

## Notes

- Administrator privileges may be required to run the program.
- Use BPF filters to exclude unnecessary packets.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

