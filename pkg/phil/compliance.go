package phil

type ComplianceRequestedEvent struct {
	RemitoID string
}

type complianceRequestedObserver struct {
	subject Subject
	readCh  chan interface{}
}

func NewComplianceRequestedObserver(subject Subject) Observer[ComplianceRequestedEvent] {
	c := &complianceRequestedObserver{
		subject: subject,
		readCh:  make(chan interface{}),
	}

	RegisterObserver(c.subject, c.readCh, c.Notify)

	return c

}

func (c complianceRequestedObserver) Close() {
	close(c.readCh)
}

func (c complianceRequestedObserver) Notify(event ComplianceRequestedEvent) {
	println("event:", event.RemitoID)
}

func DoService(subject Subject) {
	subject.Notify(ComplianceRequestedEvent{
		RemitoID: "123",
	})
}
