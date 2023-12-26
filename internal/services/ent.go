package services

import (
	"sync"
	"tests/internal/models"
	r "tests/internal/repository"
)

type EntTestService struct {
	rep *r.Repository
}

func NewEntTestService(rep *r.Repository) *EntTestService {
	return &EntTestService{
		rep,
	}
}

//	func (t *EntTestService) GetTest(mathLitID, kazHistoryID, readingLit, profile1, profile2 int) (EntTestOutput, error) {
//		output := EntTestOutput{
//			Tests: make(map[string][]models.Question),
//		}
//		var wg sync.WaitGroup
//		var mu sync.Mutex
//		errCh := make(chan error, 4)
//		wg.Add(5)
//
//		go t.asyncTestCollect(mathLitID, &wg, &mu, errCh, t.rep.Ent.GetMathLiteracyTest, &output, "mathLiteracy")
//		go t.asyncTestCollect(kazHistoryID, &wg, &mu, errCh, t.rep.Ent.GetKazHistoryTest, &output, "kazHistory")
//		go t.asyncTestCollect(readingLit, &wg, &mu, errCh, t.rep.Ent.GetReadingLiteracyTest, &output, "readingLiteracy")
//		go t.asyncTestCollect(profile1, &wg, &mu, errCh, t.rep.Ent.GetProfileTest, &output, "profile1")
//		go t.asyncTestCollect(profile2, &wg, &mu, errCh, t.rep.Ent.GetProfileTest, &output, "profile2")
//
//		go func() {
//			wg.Wait()
//			close(errCh)
//		}()
//		for err := range errCh {
//			if err != nil {
//				return EntTestOutput{}, err
//			}
//		}
//		return output, nil
//	}
//func (e *EntTestService) asyncTestCollect(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, method func(args ...interface{}) ([]models.Question, error), output *EntTestOutput, fieldName string) {
//	defer wg.Done()
//	test, err := method(subID)
//	if err != nil {
//		select {
//		case errCh <- err:
//		default:
//		}
//		return
//	}
//
//	mu.Lock()
//	output.Tests[fieldName] = test
//	mu.Unlock()
//	return
//}

func (e *EntTestService) GetTest(mathLitID, kazHistoryID, readingLit, profile1, profile2 int) (EntTestOutput, error) {
	output := EntTestOutput{
		Tests: make(map[string][]models.Question),
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 4)
	wg.Add(5)

	go func(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, output *EntTestOutput) {
		defer wg.Done()
		mathLitQuestions, err := e.buildGeneralTest(subID, 5, 0, 1, 3, 0, 2, 2, 0, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		mu.Lock()
		output.Tests["math_literacy"] = mathLitQuestions
		mu.Unlock()
	}(mathLitID, &wg, &mu, errCh, &output)

	go func(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, output *EntTestOutput) {
		defer wg.Done()
		ReadingLitQuestions, err := e.buildGeneralTest(subID, 1, 2, 1, 1, 2, 1, 1, 2, 1)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		mu.Lock()
		output.Tests["reading_literacy"] = ReadingLitQuestions
		mu.Unlock()
	}(readingLit, &wg, &mu, errCh, &output)

	go func(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, output *EntTestOutput) {
		defer wg.Done()
		KazHistoryQuestionsType0, err := e.buildGeneralTest(subID, 5, 0, 1, 3, 0, 2, 2, 0, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		KazHistoryQuestionsType2, err := e.buildGeneralTest(subID, 2, 2, 1, 0, 0, 0, 0, 0, 0)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		questions := append(KazHistoryQuestionsType0, KazHistoryQuestionsType2...)
		mu.Lock()
		output.Tests["kaz_history"] = questions
		mu.Unlock()
	}(kazHistoryID, &wg, &mu, errCh, &output)

	go func(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, output *EntTestOutput) {
		defer wg.Done()
		Profile1QuestionsType0, err := e.buildGeneralTest(subID, 15, 0, 1, 5, 0, 2, 5, 0, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		Profile1QuestionsType1, err := e.buildGeneralTest(subID, 5, 1, 1, 3, 1, 2, 2, 1, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		Profile1QuestionsType2, err := e.buildGeneralTest(subID, 1, 2, 1, 0, 0, 0, 0, 0, 0)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		questions := append(append(Profile1QuestionsType0, Profile1QuestionsType1...), Profile1QuestionsType2...)
		mu.Lock()
		output.Tests["profile_1"] = questions
		mu.Unlock()
	}(profile1, &wg, &mu, errCh, &output)

	go func(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, output *EntTestOutput) {
		defer wg.Done()
		Profile2QuestionsType0, err := e.buildGeneralTest(subID, 15, 0, 1, 5, 0, 2, 5, 0, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		Profile2QuestionsType1, err := e.buildGeneralTest(subID, 5, 1, 1, 3, 1, 2, 2, 1, 3)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		Profile2QuestionsType2, err := e.buildGeneralTest(subID, 1, 2, 1, 0, 0, 0, 0, 0, 0)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
			return
		}
		questions := append(append(Profile2QuestionsType0, Profile2QuestionsType1...), Profile2QuestionsType2...)
		mu.Lock()
		output.Tests["profile_2"] = questions
		mu.Unlock()
	}(profile2, &wg, &mu, errCh, &output)

	go func() {
		wg.Wait()
		close(errCh)
	}()
	for err := range errCh {
		if err != nil {
			return EntTestOutput{}, err
		}
	}
	return output, nil
}

func (e *EntTestService) buildGeneralTest(
	subID int, group1count, group1type, group1level int, group2count, group2type, group2level int, group3count, group3type, group3level int,
) ([]models.Question, error) {
	var questions []models.Question
	var err error
	if group1count != 0 {
		q, err := e.rep.Ent.GetQuestions(subID, group1count, group1type, group1level)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q...)
	}
	if group2count != 0 {
		q, err := e.rep.Ent.GetQuestions(subID, group2count, group2type, group2level)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q...)
	}
	if group3count != 0 {
		q, err := e.rep.Ent.GetQuestions(subID, group3count, group3type, group3level)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q...)
	}
	return questions, err
}

//func (e *EntTestService) buildProfileTest(subID int) ([]models.Question, error) {
//	var questions []models.Question
//
//}
