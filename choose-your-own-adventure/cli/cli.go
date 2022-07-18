package cli

import (
	"choose-your-own-adventure/story"
	"fmt"
)

func ShowArc(arcname string) string {
    // Get arc
    var st story.Story
	st.Load("./gopher.json")
    arc, ok := st.GetArc(arcname)
    if !ok {
        arc, _ = st.GetArc("intro")
    }
    // Show story
    fmt.Print("\n\n")
    for _, paragraph := range arc.Story {
        fmt.Println(paragraph)
    }

    // If the route ended
    if len(arc.Options) == 0 {
        return ""
    }

    // Show options
    fmt.Println("\nWhat will you do?")
    for i, option := range arc.Options {
        fmt.Println(i, option.Text)
    }

    // Get option
    var option int
    for {
        fmt.Print("Choose an option: ")
        _, err := fmt.Scanln(&option)
        // Check for valid option
        if err != nil {
            fmt.Println("You entered something not valid: ", err)
            continue
        } else if option < 0 || option >= len(arc.Options) {
            fmt.Println("The option is not correct")
        } else {
            break
        }
    }

    next := arc.Options[option].Arc
    return next
}

func Start() {
    arcname := "intro"
    for arcname != "" {
        arcname = ShowArc(arcname)
    }
}
