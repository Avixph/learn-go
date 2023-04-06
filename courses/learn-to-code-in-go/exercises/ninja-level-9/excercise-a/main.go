package main

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
