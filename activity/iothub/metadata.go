package iothub

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	AZURE_IOTHUB_CONNECTION_STRING string `md:"iothub_connection_string,required"`
	Operation string `md:operation,required`
}

type Input struct {
	DeviceId string `md:device_id,required`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error
	r.DeviceId, err = coerce.ToString(values["deviceId"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"deviceId": r.DeviceId,
	}
}

type Output struct {
	Result string `md:result`
	Status string `md:status`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	result, _ := coerce.ToObject(values["Result"])
	o.Result = result
	status, _ := coerce.ToObject(values["Status"])
	o.Status = status
	return nil
}

func (o *Output) ToMap() map[string]interface{} {

	
	return map[string]interface{}{
		"Result": o.Result,
		"Status": o.Status,
	}
}
