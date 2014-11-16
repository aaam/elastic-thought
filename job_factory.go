package elasticthought

import (
	"fmt"

	"github.com/couchbaselabs/logg"
)

// Create a new job based on the Job Descriptor
func CreateJob(config Configuration, jobDescriptor JobDescriptor) (Runnable, error) {

	// Connect to DB
	db := config.DbConnection()

	// Fetch doc associated w/ job descriptor
	doc := &ElasticThoughtDoc{}
	err := db.Retrieve(jobDescriptor.DocIdToProcess, doc)
	if err != nil {
		return nil, err
	}

	// based on document type, create the correct Runnable
	switch doc.Type {
	case DOC_TYPE_DATASET:

		// create a Dataset doc
		dataset := &Dataset{}
		err = db.Retrieve(doc.Id, &dataset)
		if err != nil {
			errMsg := fmt.Errorf("Didn't retrieve: %v - %v", doc.Id, err)
			logg.LogError(errMsg)
			return nil, errMsg
		}

		return DatasetSplitter{
			Configuration: config,
			Dataset:       *dataset,
		}, nil
	}

	return nil, fmt.Errorf("Unable to create job for: %+v", jobDescriptor)

}