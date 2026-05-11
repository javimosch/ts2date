package main
import ("fmt";"os";"strconv";"time")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: ts2date <unix_timestamp>"); os.Exit(1) }
	ts, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil { fmt.Fprintln(os.Stderr,"Invalid timestamp"); os.Exit(1) }
	fmt.Println(time.Unix(ts, 0).Format(time.RFC3339))
}
