## Should include a request validation routine 

```sh
>newanimal p1 cow
Created it!
>newanimal p2 bird
Created it!
>newanimal p3 snake
Created it!
>newanimal p4 cat
Created it!
>query p1 eat
grass
>query p2 move
fly
>query p3 speak
hsss
>query p4 eat
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x492230]

goroutine 1 [running]:
main.main()
        /home/devel1/gowkspc/src/github.com/pssslearning/courseraGoC2Week4Assignment/OtherProposals/NV/animalWithInterfaces.go:106 +0x290
exit status 2
```