// https://dotnet.microsoft.com/en-us/learn/languages/fsharp-hello-world-tutorial/modify

let PrintGreeting name = printfn $"Hello {name} from F#!"
PrintGreeting "junwei"

[1..100] |> List.sum |> printfn "sum=%d"
