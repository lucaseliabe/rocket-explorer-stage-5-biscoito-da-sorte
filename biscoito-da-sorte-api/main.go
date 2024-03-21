package main

import (
	"net/http"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type FortuneCookie struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

var fortuneCookies = []FortuneCookie{
	{1, "A vida trará coisas boas se tiveres paciência."},
	{2, "Demonstre amor e alegria em todas as oportunidades e verás que a paz nasce dentro de você."},
	{3, "Não compense na ira o que lhe falta na razão."},
	{4, "Defeitos e virtudes são apenas dois lados da mesma moeda."},
	{5, "A maior de todas as torres começa no solo."},
	{6, "Não há que ser forte. Há que ser flexível."},
	{7, "Gente todo dia arruma os cabelos, por que não o coração?"},
	{8, "Há três coisas que jamais voltam; a flecha lançada, a palavra dita e a oportunidade perdida."},
	{9, "A juventude não é uma época da vida, é um estado de espírito."},
	{10, "Podemos escolher o que semear, mas somos obrigados a colher o que plantamos."},
	{11, "Dê toda a atenção para a formação dos teus filhos, sobretudo por exemplos de tua própria vida."},
	{12, "Siga os bons e aprenda com eles."},
	{13, "Não importa o tamanho da montanha, ela não pode tapar o sol."},
	{14, "O bom-senso vai mais longe do que muito conhecimento."},
	{15, "Quem quer colher rosas deve suportar os espinhos."},
	{16, "São os nossos amigos que nos ensinam as mais valiosas lições."},
	{17, "Aquele que se importa com o sentimento dos outros, não é um tolo."},
	{18, "A adversidade é um espelho que reflete o verdadeiro eu."},
	{19, "Lamentar aquilo que não temos é desperdiçar aquilo que já possuímos."},
	{20, "Uma bela flor é incompleta sem suas folhas."},
	{21, "Sem o fogo do entusiasmo, não há o calor da vitória."},
	{22, "Não há melhor negócio que a vida. A gente há obtém a troco de nada."},
	{23, "O riso é a menor distância entre duas pessoas."},
	{24, "Você é jovem apenas uma vez. Depois precisa inventar outra desculpa."},
	{25, "É mais fácil conseguir o perdão do que a permissão."},
	{26, "Acredite em si mesmo e tudo será possível."},
	{27, "A vida é curta, aproveite cada momento."},
	{28, "A persistência é o caminho do êxito."},
	{29, "A sorte favorece a mente preparada."},
	{30, "A verdadeira sabedoria está em reconhecer a própria ignorância."},
	{31, "A felicidade está nas pequenas coisas."},
	{32, "A gratidão transforma o que temos em suficiente."},
	{33, "A imaginação é mais importante que o conhecimento."},
	{34, "A paciência é uma virtude."},
	{35, "A simplicidade é o último grau de sofisticação."},
	{36, "A vida é como andar de bicicleta. Para manter o equilíbrio, você precisa se manter em movimento."},
	{37, "Acredite em milagres, mas não dependa deles."},
	{38, "Amar é encontrar na felicidade de outrem a própria felicidade."},
	{39, "Aprenda como se você fosse viver para sempre."},
	{40, "As melhores coisas da vida não são coisas."},
	{41, "As pessoas mais felizes não têm as melhores coisas, elas sabem fazer o melhor das oportunidades que aparecem em seus caminhos."},
	{42, "Cada dia é uma nova chance para ser feliz."},
	{43, "Deixe de lado as preocupações e aproveite a vida."},
	{44, "Dê o seu melhor e o universo fará o resto."},
	{45, "Difícil é ganhar um amigo em uma hora; fácil é ofendê-lo em um minuto."},
	{46, "É melhor ser odiado pelo que você é do que ser amado pelo que você não é."},
	{47, "Encontre a alegria nas pequenas coisas."},
	{48, "Faça o que você pode, com o que você tem, onde você está."},
	{49, "Felicidade é a combinação de sorte com escolhas bem feitas."},
	{50, "Lembre-se de que a felicidade é uma viagem, não um destino."},
	{51, "Mantenha o foco no objetivo e siga em frente."},
	{52, "Nada é impossível para aquele que persiste."},
	{53, "Nunca é tarde demais para ser o que você poderia ter sido."},
	{54, "O futuro pertence àqueles que acreditam na beleza de seus sonhos."},
	{55, "O importante não é vencer todos os dias, mas lutar sempre."},
	{56, "O sucesso é a soma de pequenos esforços repetidos dia após dia."},
	{57, "O verdadeiro segredo da felicidade está em ter um genuíno interesse por todos os detalhes da vida."},
	{58, "Pare de esperar que as coisas aconteçam. Vá lá e faça acontecer."},
	{59, "Seja a mudança que você deseja ver no mundo."},
	{60, "Seja feliz com o que você tem enquanto trabalha pelo que você quer."},
}

func main() {
	router := gin.Default()

	router.GET("/fortunecookies", getFortuneCookies)
	router.GET("/fortunecookies/random", getRandFortuneCookie)

	router.Run(":8080")
}

func getFortuneCookies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fortuneCookies)
}

func getRandFortuneCookie(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fortuneCookies[rand.Intn(len(fortuneCookies))])
}