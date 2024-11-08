# G.I. Go - Générateurs & Itérateurs en Go

<img alt="Illustration de gentil G.I. Gopher" src="./images/gi%20gopher.png" width="200px" /> <!-- markdownlint-disable-line MD033: inline HTML required for scaling images -->

_(illustration générée avec ChatGPT)_ <!-- markdownlint-disable-line MD036: this is italics on purpose -->

## ou « Comment compter les Gophers sans perdre la mémoire »

[![MIT license](https://img.shields.io/badge/license-MIT-green)](LICENSE)

Ce dépôt rassemble les exemples de code et les slides de ma conférence sur les itérateurs Go, introduits en 2024 dans la [version 1.23](https://go.dev/doc/go1.23#iterators).

### Contenu

Il y a 3 grandes parties

1. [`main1.go`](./main1.go) et [`backward.go`](./backward.go)

   Illustration du mécanisme de base, en parcourant un tableau à l'envers à l'aide d'un générateur [`iter.Seq`](https://pkg.go.dev/iter#Seq) puis d'un itérateur [`iter.Pull`](https://pkg.go.dev/iter#Pull).

2. [`main2.go`](./main2.go) et [`words.go`](./words.go)

   Application à un générateur de mots issus d'un fichier texte, et comparaison des performances avec la méthode naïve, ainsi qu'avec l'utilisation d'un channel.

3. [`main3.go`](./main3.go) et [`count.go`](./count.go)

   Application au décompte des mots _distincts_ d'un fichier, de façon exacte sans limitation, ou approximativement dans un espace mémoire contraint, à l'aide de l'[algorithme probabiliste CVM](https://www.quantamagazine.org/computer-scientists-invent-an-efficient-new-way-to-count-20240516/).

### Présentations

Cette conférence a été présentée :

- au [meetup Golang Rennes](https://www.meetup.com/fr-FR/golang-rennes/events/303884251/) le 19 novembre 2024
