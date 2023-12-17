package services

import (
	"fmt"
	"sync"
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

func (t *EntTestService) Get(subjectID int) {
	res := t.rep.Ent.Get(subjectID)

	fmt.Println(res)
}

func (t *EntTestService) GetTest(mathLitID, kazHistoryID, readingLit, profile1, profile2 int) (EntTestOutput, error) {
	var output EntTestOutput
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 4)
	wg.Add(4)

	go t.asyncTestCollect(mathLitID, &wg, &mu, errCh, t.rep.Ent.GetMathLiteracyTest, &output, "mathLiteracy")
	go t.asyncTestCollect(kazHistoryID, &wg, &mu, errCh, t.rep.Ent.GetKazHistoryTest, &output, "kazHistory")
	go t.asyncTestCollect(readingLit, &wg, &mu, errCh, t.rep.Ent.GetReadingLiteracyTest, &output, "readingLiteracy")
	go t.asyncTestCollect(profile1, &wg, &mu, errCh, t.rep.Ent.GetProfileTest, &output, "profile1")
	go t.asyncTestCollect(profile2, &wg, &mu, errCh, t.rep.Ent.GetProfileTest, &output, "profile2")

	go func() {
		wg.Wait()
		close(errCh)
	}()
	for err := range errCh {
		if err != nil {
			return EntTestOutput{}, err
		}
	}
	wg.Wait()
	return output, nil
}

func (t *EntTestService) asyncTestCollect(subID int, wg *sync.WaitGroup, mu *sync.Mutex, errCh chan error, method func(int) (r.TestOutput, error), output *EntTestOutput, fieldName string) {
	defer wg.Done()
	test, err := method(subID)
	if err != nil {
		select {
		case errCh <- err:
		default:
		}
		return
	}

	mu.Lock()
	output.Tests[fieldName] = test
	mu.Unlock()
	return
}
