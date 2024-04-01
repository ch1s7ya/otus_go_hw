package hw03frequencyanalysis

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestGetTop10(t *testing.T) {
	type args struct {
		words []WordFrequency
	}
	tests := []struct {
		name string
		args args
		want []WordFrequency
	}{
		{
			name: "Small slice",
			args: args{
				words: []WordFrequency{
					{
						Word:  "hello",
						Count: 1,
					},
					{
						Word:  "world",
						Count: 1,
					},
				},
			},
			want: []WordFrequency{
				{
					Word:  "hello",
					Count: 1,
				},
				{
					Word:  "world",
					Count: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTop10(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTop10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildFrequencyListOfWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Regular words",
			args: args{
				words: []string{"bite", "my", "shiny", "metal", "ass"},
			},
			want: map[string]int{"bite": 1, "my": 1, "shiny": 1, "metal": 1, "ass": 1},
		},
		{
			name: "The same number of words",
			args: args{
				words: []string{"shut", "up", "and", "and", "money", "take", "my", "money"},
			},
			want: map[string]int{"shut": 1, "up": 1, "and": 2, "take": 1, "my": 1, "money": 2},
		},
		{
			name: "Empty text",
			args: args{
				words: []string{},
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildFrequencyListOfWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildFrequencyListOfWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLowerAndTrim(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty array",
			args: args{
				words: []string{},
			},
			want: nil,
		},
		{
			name: "Regular words",
			args: args{
				words: []string{"Shut", "up", "and", "take", "my", "money!"},
			},
			want: []string{"shut", "up", "and", "take", "my", "money"},
		},
		{
			name: "Strange words",
			args: args{
				words: []string{"-------", "-", "dog,cat", "dog...cat", "dogcat"},
			},
			want: []string{"-------", "dog,cat", "dog...cat", "dogcat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerAndTrim(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLowerAndTrim() = %v, want %v", got, tt.want)
			}
		})
	}
}
