package service

import (
	"reflect"
	"testing"
	word_pb "wcservice/grpc/pb/word"
)

func Test_wordService_GetWords10(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		s    *wordService
		args args
		want []*word_pb.WordCount
	}{
		{
			"Success",
			nil,
			args{
				text: "ab cd ef ab ab ef",
			},
			[]*word_pb.WordCount{
				{
					Word:  "ab",
					Count: 3,
				},
				{
					Word:  "ef",
					Count: 2,
				},
				{
					Word:  "cd",
					Count: 1,
				},
			},
		},
		{
			"Success 10000",
			nil,
			args{
				text: `a
				aa
				aa
				aaa
				aaa
				aaa
				aaron
				aaron
				aaron
				aaron
				ab
				ab
				ab
				ab
				ab
				abandoned
				abandoned
				abandoned
				abandoned
				abandoned
				abandoned
				abc
				abc
				abc
				abc
				abc
				abc
				abc
				aberdeen
				aberdeen
				aberdeen
				aberdeen
				aberdeen
				aberdeen
				aberdeen
				aberdeen
				abilities
				abilities
				abilities
				abilities
				abilities
				abilities
				abilities
				abilities
				abilities
				ability
				ability
				ability
				ability
				ability
				ability
				ability
				ability
				ability
				ability
				able
				able
				able
				able
				able
				able
				able
				able
				able
				able
				able
				abuse
				acc
				accent
				accept
				acceptable
				acceptance
				accepted
				accepting
				accessing
				accessories
				accurately
				accused
				acdbentity
				ace
				acer
				achieve
				achieved
				achievement
				achievements
				achieving
				acid
				acids
				acknowledge
				acknowledged
				acm
				acne
				acoustic
				acquire
				acquired
				acquisition
				acquisitions
				acre
				acres
				acrobat
				across
				acrylic
				act
				acting
				action
				actions
				activated
				activation
				active
				actively
				activists
				activities
				activity
				actor
				actors`,
			},
			[]*word_pb.WordCount{
				{
					Word:  "able",
					Count: 11,
				},
				{
					Word:  "ability",
					Count: 10,
				},
				{
					Word:  "abilities",
					Count: 9,
				},
				{
					Word:  "aberdeen",
					Count: 8,
				},
				{
					Word:  "abc",
					Count: 7,
				},
				{
					Word:  "abandoned",
					Count: 6,
				},
				{
					Word:  "ab",
					Count: 5,
				},
				{
					Word:  "aaron",
					Count: 4,
				},
				{
					Word:  "aaa",
					Count: 3,
				},
				{
					Word:  "aa",
					Count: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := WordServiceN_()
			if got := s.GetWords10(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordService.GetWords10() = %v, want %v", got, tt.want)
			}
		})
	}
}
