#let project(title: "", authors: (), date: none, body) = {
  // Set the document's basic properties.
  set document(author: authors, title: title)

  set page(
    paper: "a4",
    margin: (left: 3cm, right: 1.5cm, top: 2cm, bottom: 1.45cm),
    numbering: "1",
    number-align: top+center,
  )

  // Save the font familie in variable.
  let main-font = "Times New Roman MT Std"


  // Set body font family.
  set text(font: main-font, size: 14pt, lang: "de")


  // Contents
  show outline.entry.where(
    level: 1
  ): it => {
    v(19pt, weak: true)
    it
  }

  show outline.entry.where(
    level: 2
  ): it => {
    v(19pt, weak: true)
    it
  }

  show outline: it =>{align(center)[#it]}

  // Links
  show link: it => [
    #set text(rgb("#0000ff"))
    #it
  ]


  // Raw text in parapraph must have main font
  show raw.where(block: false): it => {text(font: main-font, size: 14pt)[#it]}



  //-------------------------Title-page------------------------------//
  
  set page(
    numbering: none,
    header: [
      #align(center)[

        #block(text(
          font: main-font, 
          weight: "semibold", 
          14pt, 
          "Ministerium für Bildung und Wissenschaft der Kirgisischen Republik\n" + 
          ""
        ))
      ]
    ],
    footer: [
      #align(center)[
        #block(text(
          font: main-font, 
          weight: "semibold", 
          13pt, 
          "Bischkek 2024" 
        ))
      ]
    ]
  )

  align(center)[
    #block(text(
      font: main-font, 
      weight: "semibold", 
      14pt, 
      "KIRGISISCH-DEUTSCHES INSTITUT FÜR ANGEWANDTE INFORMATIK",
    ))
  ]



  // Vertical skip
  v(12em)

  // Title row.
  align(center)[
    #block(text(font: main-font, weight: "bold", 30pt, "SEMESTERARBEIT"))

    #v(0.5em)

    #block(text(font: main-font, weight: "regular", 18pt, "zum Thema: " + "“"+title+"”"))


    #v(1.5em, weak: true)


    // #date
    #v(8em, weak: true)
  ]




  
  // Author information.
    let auths = authors.map(author => align(right, 
      text(
        weight: "medium",
        14pt,
        author + "_________" 
      )))


    align(right,
    text(
      weight: "semibold",
       14pt,
      "Erstellt von:\n" 
    )+
    text(
      weight: "medium",
       14pt,
      "\nStudierende Gruppe ABC-1-23:\n" 
    ),
  )
    auths.join(v(1pt), )


    align(right,
      text(
        weight: "semibold",
         14pt,
        "\n Betreuer: "
      )+
      text(
        weight: "medium",
         14pt,
        "\n\nBergamot P.I._________\n\n" + 
        "Natrinomatov R.T._________\n\n"
      )+ 
      text(
        weight: "semibold",
         14pt,
        "Berater für deutsche Sprache:\n\n"
      )+ 
      text(
        weight: "medium",
         14pt,
        "Nitaratieva A.Т._________\n" + v(1pt) +
        "Sarkymbaevava A.Zh._________\n" + v(1pt)
      )
    )


  //-----------------------------------------------------------------//




  //-------------------------The-Outline-----------------------------//
  // we don't need previously defined header and footer in Contents pages
  set page(
    header: [
      #counter(page).display(
            "1",
      )
  ],
    footer: [
  ],
  )

  // Table of contents
  show heading: it => {text(weight: "regular", font: main-font)[#it]}
  show outline: it => {text(weight: "regular", font: main-font)[#it]}
  outline(
    depth: 2,
  )

  // Rest of the Headings must be bold
  show heading: it => {strong(it)}
  //-----------------------------------------------------------------//



  // Header and Footer 
  // They are here cause we don't need them in the title page
  set page(
    numbering: "1",
    // // Footer of the page
    header: [
      #align(center)[
        #counter(page).display("1")
      ]
    ],
    )


  // Parapraph
  set par(justify: true, first-line-indent: 1.25cm, leading: 20pt)

  // Even first parapraph must have first line(red line) indent
  show heading: it =>  {
    it
    par()[#text(size:0.5em)[#h(0.0em)]]
  }


  // Even first parapraph must have first line(red line) indent
  show heading.where(level: 3): it =>  {
    v(20pt, weak: true)
    it
    par()[#text(size:0.5em)[#h(0.0em)]]
  }
  // Even first parapraph must have first line(red line) indent
  show heading.where(level: 2): it =>  {
    v(20pt, weak: true)
    it
    par()[#text(size:0.5em)[#h(0.0em)]]
  }

  // New chapter starts with new page
  show heading.where(level: 1): it => {
    pagebreak()
    it
    par()[#text(size:0.5em)[#h(0.0em)]]
  }
  

  show heading: set block(above: 0.7em, below: 0.7em)

  // Heading numbering
  show heading.where(level: 1, outlined: true, numbering: "1."): set heading(
    numbering: 
      (..nums) => "Kapitel " + nums
      .pos()
      .map(str)
      .join(".") + "."
  )

  show heading.where(outlined: true): set heading(numbering: "1.") 

  // Heading centering
  show heading: it => {align(center)[#it]}



  body
}
