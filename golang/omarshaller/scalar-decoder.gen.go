// Code generated for marshaling/unmarshaling. DO NOT EDIT.
package omarshaller

import (
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
	"reflect"
)


func (d *Decoder) getInt8(number interface{})(value int8, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or int8), got nil/null")
		return
	}

    value, ok = number.(int8)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or int8), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinInt8 <= fv && fv <= hybrids.MaxInt8){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = int8(fv)
	return
}

func (d *Decoder) decodeInt8(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Int8WriterAccessor) (err error) {
    var value int8

    value, err = d.getInt8(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int8,
          OmniqlType:  "Int8",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetInt8(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int8,
          OmniqlType:  "Int8",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getUint8(number interface{})(value uint8, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or uint8), got nil/null")
		return
	}

    value, ok = number.(uint8)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or uint8), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinUint8 <= fv && fv <= hybrids.MaxUint8){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = uint8(fv)
	return
}

func (d *Decoder) decodeUint8(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Uint8WriterAccessor) (err error) {
    var value uint8

    value, err = d.getUint8(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint8,
          OmniqlType:  "Uint8",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetUint8(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint8,
          OmniqlType:  "Uint8",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getInt16(number interface{})(value int16, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or int16), got nil/null")
		return
	}

    value, ok = number.(int16)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or int16), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinInt16 <= fv && fv <= hybrids.MaxInt16){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = int16(fv)
	return
}

func (d *Decoder) decodeInt16(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Int16WriterAccessor) (err error) {
    var value int16

    value, err = d.getInt16(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int16,
          OmniqlType:  "Int16",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetInt16(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int16,
          OmniqlType:  "Int16",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getUint16(number interface{})(value uint16, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or uint16), got nil/null")
		return
	}

    value, ok = number.(uint16)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or uint16), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinUint16 <= fv && fv <= hybrids.MaxUint16){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = uint16(fv)
	return
}

func (d *Decoder) decodeUint16(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Uint16WriterAccessor) (err error) {
    var value uint16

    value, err = d.getUint16(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint16,
          OmniqlType:  "Uint16",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetUint16(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint16,
          OmniqlType:  "Uint16",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getInt32(number interface{})(value int32, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or int32), got nil/null")
		return
	}

    value, ok = number.(int32)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or int32), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinInt32 <= fv && fv <= hybrids.MaxInt32){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = int32(fv)
	return
}

func (d *Decoder) decodeInt32(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Int32WriterAccessor) (err error) {
    var value int32

    value, err = d.getInt32(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int32,
          OmniqlType:  "Int32",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetInt32(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Int32,
          OmniqlType:  "Int32",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getUint32(number interface{})(value uint32, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or uint32), got nil/null")
		return
	}

    value, ok = number.(uint32)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or uint32), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinUint32 <= fv && fv <= hybrids.MaxUint32){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = uint32(fv)
	return
}

func (d *Decoder) decodeUint32(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Uint32WriterAccessor) (err error) {
    var value uint32

    value, err = d.getUint32(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint32,
          OmniqlType:  "Uint32",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetUint32(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Uint32,
          OmniqlType:  "Uint32",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}


func (d *Decoder) getFloat32(number interface{})(value float32, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("number expected (float64 or float32), got nil/null")
		return
	}

    value, ok = number.(float32)
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("number expected (float64 or float32), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.MinFloat32 <= fv && fv <= hybrids.MaxFloat32){
        err = fmt.Errorf("number is out of bound, got %d", fv)
		return
	}

	value = float32(fv)
	return
}

func (d *Decoder) decodeFloat32(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Float32WriterAccessor) (err error) {
    var value float32

    value, err = d.getFloat32(number)

    if err!=nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Float32,
          OmniqlType:  "Float32",
          ErrorMsg:    err.Error(),
          }
       return
    }

    err = tw.SetFloat32(fn, value)

    if err != nil {
       err = &DecodeError{
          Path:        path,
          Application: d.application,
          HybridType:  hybrids.Float32,
          OmniqlType:  "Float32",
          ErrorMsg:    err.Error(),
       }
       return
    }
    return
}

func (d *Decoder) decodeVectorInt8(path string, value interface{}, vw hybrids.VectorInt8Writer) (err error) {
    var item int8
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorInt8,
			OmniqlType:  "Vector[Int8]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getInt8(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorInt8,
               OmniqlType:  "Vector[Int8]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushInt8(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorInt8,
               OmniqlType:  "Vector[Int8]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorUint8(path string, value interface{}, vw hybrids.VectorUint8Writer) (err error) {
    var item uint8
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorUint8,
			OmniqlType:  "Vector[Uint8]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getUint8(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorUint8,
               OmniqlType:  "Vector[Uint8]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushUint8(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorUint8,
               OmniqlType:  "Vector[Uint8]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorInt16(path string, value interface{}, vw hybrids.VectorInt16Writer) (err error) {
    var item int16
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorInt16,
			OmniqlType:  "Vector[Int16]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getInt16(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorInt16,
               OmniqlType:  "Vector[Int16]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushInt16(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorInt16,
               OmniqlType:  "Vector[Int16]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorUint16(path string, value interface{}, vw hybrids.VectorUint16Writer) (err error) {
    var item uint16
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorUint16,
			OmniqlType:  "Vector[Uint16]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getUint16(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorUint16,
               OmniqlType:  "Vector[Uint16]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushUint16(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorUint16,
               OmniqlType:  "Vector[Uint16]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorInt32(path string, value interface{}, vw hybrids.VectorInt32Writer) (err error) {
    var item int32
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorInt32,
			OmniqlType:  "Vector[Int32]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getInt32(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorInt32,
               OmniqlType:  "Vector[Int32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushInt32(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorInt32,
               OmniqlType:  "Vector[Int32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorUint32(path string, value interface{}, vw hybrids.VectorUint32Writer) (err error) {
    var item uint32
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorUint32,
			OmniqlType:  "Vector[Uint32]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getUint32(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorUint32,
               OmniqlType:  "Vector[Uint32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushUint32(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorUint32,
               OmniqlType:  "Vector[Uint32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorFloat32(path string, value interface{}, vw hybrids.VectorFloat32Writer) (err error) {
    var item float32
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorFloat32,
			OmniqlType:  "Vector[Float32]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getFloat32(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorFloat32,
               OmniqlType:  "Vector[Float32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushFloat32(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorFloat32,
               OmniqlType:  "Vector[Float32]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorFloat64(path string, value interface{}, vw hybrids.VectorFloat64Writer) (err error) {
    var item float64
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorFloat64,
			OmniqlType:  "Vector[Float64]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getFloat64(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorFloat64,
               OmniqlType:  "Vector[Float64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushFloat64(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorFloat64,
               OmniqlType:  "Vector[Float64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorInt64(path string, value interface{}, vw hybrids.VectorInt64Writer) (err error) {
    var item int64
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorInt64,
			OmniqlType:  "Vector[Int64]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getInt64(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorInt64,
               OmniqlType:  "Vector[Int64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushInt64(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorInt64,
               OmniqlType:  "Vector[Int64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}

func (d *Decoder) decodeVectorUint64(path string, value interface{}, vw hybrids.VectorUint64Writer) (err error) {
    var item uint64
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorUint64,
			OmniqlType:  "Vector[Uint64]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.getUint64(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.VectorUint64,
               OmniqlType:  "Vector[Uint64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.PushUint64(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.VectorUint64,
               OmniqlType:  "Vector[Uint64]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}
