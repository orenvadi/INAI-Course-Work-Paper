#import "@preview/codelst:2.0.0": sourcecode, sourcefile


// #heading(level: 1, outlined: true, numbering: "1.")[Mobile Anwendung]

#heading(
level: 1, outlined: true,
    numbering: (..nums) => "Kapitel " + nums
      .pos()
      .map(str)
      .join(".") + "."
  )[Mobile Anwendung]


// Es müssen mindestens 2 Partitionen vorhanden sein

Für Studierende ist eine mobile Anwendung erforderlich, mit der sie: ihren Stundenplan einsehen und ihre Anwesenheit im System durch Scannen eines Codes registrieren können

== Funktionen mobiler Apps


- Registrierung/Anmeldung zum Konto
- Zeitplan anzeigen
- QR-Code scannen

=== Schnittstelle für mobile Anwendungen



#grid(
  columns: (auto, auto),
  rows: (auto),

  figure(
    image("../images/Login-mobile.png", width: 45%),
    caption: [
    Loginseite
    ],
  ),

  figure(
    image("../images/Subjects-mobile.png", width: 45%),
    caption: [
    Seite „Zeitplan“
    ],
  )

)

#grid(
  columns: (auto, auto),
  rows: (auto),

  figure(
    image("../images/Settings-mobile.png", width: 45%),
    caption: [
   Seite „Einstellungen“ ],
  ),

  figure(
    image("../images/Scan-mobile.png", width: 45%),
    caption: [
QR-Code-Scanseite
    ],
  )

)




=== Verwendete Techniken

Programmiersprache für mobile @android-doc Anwendungen: Kotlin @kotlin-doc
Programmiersprache für Web Verwendung: JS 

Entwickler: Android Studio, Neovim

== Projektstruktur
#figure(
  sourcefile(read("../images/mobiletree.sh"), file:"../images/mobiletree.sh"),
  caption: [
Mobile App-Projektstruktur]
)


=== Projektarchitektur

Es wurde die MVI-Architektur verwendet.
Modellansichtsabsicht

#figure(
  image("../images/mvi-arch.png", width: 85%),
  caption: [
  MVI arch Schema
]
)



== Beschreibung des Quellcodes

=== Models

#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/model/AuditoryExercise.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/model/AuditoryExercise.kt")
#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Student.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Student.kt")
#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Subject.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Subject.kt")
#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Lecture.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/model/Lecture.kt")
#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/model/IActivity.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/model/IActivity.kt")

#align(center)[Listing 2: Android App models]

- IActivity:- Diese Schnittstelle definiert zwei Eigenschaften: `activityType` (String) und `activityDuration` (Int). Sie dient als Vorlage für andere Datenklassen, die Aktivitäten darstellen.
- AuditoryExercise:- Diese Klasse erbt von `IActivity` und stellt eine Hörübung innerhalb einer Vorlesung dar. Sie hat zwei Eigenschaften: `activityType` (String) und `activityDuration` (Int), die von der Schnittstelle geerbt wurden und wahrscheinlich die Art der Übung und ihre Dauer angeben.
- Lecture:- Diese Klasse stellt eine Vorlesung dar und hat die folgenden Eigenschaften:
- `id` (String?): Eindeutige Kennung für die Vorlesung (möglicherweise null).
- `date` (String?): Datum der Vorlesung (möglicherweise null).
- `type` (String): Art der Vorlesung (z. B. Seminar, Kurs).
- `number` (Int?): Nummer der Vorlesung in einem Kurs (möglicherweise null).
- `professorId` (String?): Eindeutige Kennung für den Professor.
- `studentIds` (MutableList<String>): Liste der Studenten-IDs, die an der Vorlesung teilnehmen.
- Student:- Diese Klasse repräsentiert einen Studenten und hat die folgenden Eigenschaften:
- `id` (String?): Eindeutige Kennung für den Studenten (möglicherweise null).
- `name` (String?): Name des Studenten (möglicherweise null).
- `email` (String?): E-Mail-Adresse des Studenten (möglicherweise null).


=== MainActivity

#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/MainActivity.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/MainActivity.kt")
#align(center)[Listing 3: MainActivity.java]


*Importe:*

- Standard-Android-Bibliotheken zum Arbeiten mit Aktivitäten, Fragmenten und Layouts. 
- Firebase-Bibliotheken für Benutzerauthentifizierung und Zugriff auf Firestore-Datenbanken.

*Klasse:*

- `MainActivity` erweitert `AppCompatActivity`, die Basisklasse für die meisten Aktivitäten in Android.

*Mitgliedsvariablen:*

- `auth`: Instanz von `FirebaseAuth`, die für die Benutzerauthentifizierung mit Firebase verwendet wird.
- `db`: Instanz von `FirebaseFirestore`, die für den Zugriff auf die Firestore-Datenbank verwendet wird.
- `fragmentManager`: Verweis auf den `FragmentManager`, der zum Verwalten von Fragmenten innerhalb der Aktivität verwendet wird.

*onCreate-Methode:*

- Diese Methode wird aufgerufen, wenn die Aktivität zum ersten Mal erstellt wird.
- Legt das Layout der Aktivität mit `setContentView` fest.
- Sucht die `FragmentContainerView` anhand ihrer ID (`R.id.fragmentContainerView`). Dieser Container enthält verschiedene Fragmente innerhalb der Aktivität.
- Findet die Textansicht mit der ID „app_name_title“.
- Legt einen onClickListener für die Textansicht „appTitle“ fest.
- Überprüft, ob das aktuelle Fragment, das in der „fragmentContainerView“ angezeigt wird, kein „WelcomeFragment“ ist.
- Wenn es nicht der Willkommensbildschirm ist, wird eine neue Fragmenttransaktion erstellt und das aktuelle Fragment durch ein „WelcomeFragment“ ersetzt.
- Erstellt eine neue Fragmenttransaktion und ersetzt den Inhalt der „fragmentContainerView“ durch ein neues „WelcomeFragment“. Dadurch wird sichergestellt, dass die App auf dem Willkommensbildschirm startet.

*onBackPressed-Methode:*

- Diese Methode wird aufgerufen, wenn der Benutzer die Zurück-Taste auf dem Gerät drückt.
- Überprüft, ob das aktuelle Fragment ein „WelcomeFragment“ ist.
- Wenn es der Willkommensbildschirm ist, wird die App vollständig mit „finishAffinity“ beendet.
- Wenn es nicht der Willkommensbildschirm ist, wird der Back-Stack des Fragment-Managers geöffnet. Dies navigiert im Wesentlichen zurück zum vorherigen Fragment.



=== StudentRecyclerAdapter


#sourcefile(read("../../AndroidApp/java/com/marko112/lecturepresenceapp/student/StudentRecyclerAdapter.kt"), file:"../../AndroidApp/java/com/marko112/lecturepresenceapp/student/StudentRecyclerAdapter.kt")
#align(center)[Listing 4: StudentRecyclerAdapter.java]

Dieser Code definiert eine `StudentRecyclerAdapter`-Klasse, die mit einem RecyclerView in einer Android-Anwendung verwendet wird. Hier ist eine Aufschlüsselung:

*Paket:*

- `com.marko112.lecturepresenceapp.student`: Dies zeigt an, dass der Adapter wahrscheinlich spezifisch für Studentendaten innerhalb der Vorlesungspräsenzanwendung ist.

*Klasse:*

- `StudentRecyclerAdapter` erweitert `RecyclerView.Adapter`. Diese Klasse verwaltet, wie Daten in einem RecyclerView angezeigt werden.

*Mitgliedsvariablen:*

- `items`: Eine veränderbare Liste von `Student`-Objekten. Diese Liste enthält die Daten, die im RecyclerView angezeigt werden sollen.

*Konstruktor:*

- Der Konstruktor verwendet eine `MutableList<Student>` als Eingabe und weist sie der Variable `items` zu.

*onCreateViewHolder-Methode:*

- Diese Methode wird vom RecyclerView aufgerufen, wenn eine neue Ansicht benötigt wird.
- Es bläst das Layout für ein einzelnes Studentenelement (definiert in `student_recycler_item.xml`) mithilfe des `LayoutInflater` auf.

- Es erstellt eine neue `StudentViewHolder`-Instanz und übergibt die aufgeblähte Ansicht an diese.

