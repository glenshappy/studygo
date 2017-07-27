    package main
    import (  
	"runtime/pprof"
        "os"
	"fmt" 
    )  
    func sum(j int,c chan int){
	c <-j
    }
    func main() {  
	var maxnum int=10;
	c:=make(chan int,maxnum)
	defer close(c)
        f, _ := os.Create("profile_file")  
        pprof.StartCPUProfile(f)  // 开始cpu profile，结果写到文件f中  
        defer pprof.StopCPUProfile()  // 结束profile  
        for i:=0;i<maxnum;i++ {
		go sum(i,c);
	}  
	
	for a,ok:= <- c;ok!=false; {
	    fmt.Println("from channel c:",a);
	}
	fmt.Println("Done!");
    }
