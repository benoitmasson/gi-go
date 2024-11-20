# G.I. Go - Générateurs & Itérateurs en Go

<img alt="Illustration de gentil G.I. Gopher" src="./slides/images/gi%20gopher.png" width="200px" /> <!-- markdownlint-disable-line MD033: inline HTML required for scaling images -->

_(illustration générée avec [ChatGPT](https://chatgpt.com/))_ <!-- markdownlint-disable-line MD036: this is italics on purpose -->

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
  🔎 [_slides_](../golang-rennes/slides/gi-go.pdf)

### Sources et liens utiles

- Contexte et décisions, dans la proposition initiale : https://github.com/golang/go/discussions/56413
- Article du blog Golang qui détaille le fonctionnement des itérateurs : https://go.dev/blog/range-functions
- Discussions avec [Thibaut Rousseau](https://github.com/Thiht), notamment pour les [tests unitaires](./backward_test.go)
- Arguments contre l'ajout des itérateurs : https://itnext.io/go-evolves-in-the-wrong-direction-7dfda8a1a620
- Publication scientifique présentant l'algorithme CVM : https://arxiv.org/pdf/2301.10191
- Article décrivant l'intérêt et les cas d'usage du nouveau package [`unique`](https://pkg.go.dev/unique), introduit également dans Go 1.23, que je n'ai pas réussi à utiliser de façon efficace [dans cette démo](https://github.com/benoitmasson/gi-go/commit/07c7482d1033e520b56cb2cb5ebac66f31ebb238) : https://medium.com/google-cloud/interning-in-go-4319ea635002
