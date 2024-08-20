package proof

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/snadrus/must"
	"github.com/stretchr/testify/require"
)

const TAuxGolden = `CwAAAAAAAAAhAAAAAAAAAC9kYXRhLzEvc3RvcmUvY2FjaGUvcy10MDI2MjAtMjAxMAcAAAAAAAAA
bGF5ZXItMQEAAABAAAAAAAAAAAAAAAAAIQAAAAAAAAAvZGF0YS8xL3N0b3JlL2NhY2hlL3MtdDAy
NjIwLTIwMTAHAAAAAAAAAGxheWVyLTIBAAAAQAAAAAAAAAAAAAAAACEAAAAAAAAAL2RhdGEvMS9z
dG9yZS9jYWNoZS9zLXQwMjYyMC0yMDEwBwAAAAAAAABsYXllci0zAQAAAEAAAAAAAAAAAAAAAAAh
AAAAAAAAAC9kYXRhLzEvc3RvcmUvY2FjaGUvcy10MDI2MjAtMjAxMAcAAAAAAAAAbGF5ZXItNAEA
AABAAAAAAAAAAAAAAAAAIQAAAAAAAAAvZGF0YS8xL3N0b3JlL2NhY2hlL3MtdDAyNjIwLTIwMTAH
AAAAAAAAAGxheWVyLTUBAAAAQAAAAAAAAAAAAAAAACEAAAAAAAAAL2RhdGEvMS9zdG9yZS9jYWNo
ZS9zLXQwMjYyMC0yMDEwBwAAAAAAAABsYXllci02AQAAAEAAAAAAAAAAAAAAAAAhAAAAAAAAAC9k
YXRhLzEvc3RvcmUvY2FjaGUvcy10MDI2MjAtMjAxMAcAAAAAAAAAbGF5ZXItNwEAAABAAAAAAAAA
AAAAAAAAIQAAAAAAAAAvZGF0YS8xL3N0b3JlL2NhY2hlL3MtdDAyNjIwLTIwMTAHAAAAAAAAAGxh
eWVyLTgBAAAAQAAAAAAAAAAAAAAAACEAAAAAAAAAL2RhdGEvMS9zdG9yZS9jYWNoZS9zLXQwMjYy
MC0yMDEwBwAAAAAAAABsYXllci05AQAAAEAAAAAAAAAAAAAAAAAhAAAAAAAAAC9kYXRhLzEvc3Rv
cmUvY2FjaGUvcy10MDI2MjAtMjAxMAgAAAAAAAAAbGF5ZXItMTABAAAAQAAAAAAAAAAAAAAAACEA
AAAAAAAAL2RhdGEvMS9zdG9yZS9jYWNoZS9zLXQwMjYyMC0yMDEwCAAAAAAAAABsYXllci0xMQEA
AABAAAAAAAAAAAAAAAAAIQAAAAAAAAAvZGF0YS8xL3N0b3JlL2NhY2hlL3MtdDAyNjIwLTIwMTAG
AAAAAAAAAHRyZWUtZAH///9/AAAAAAAAAAAAAAAAIQAAAAAAAAAvZGF0YS8xL3N0b3JlL2NhY2hl
L3MtdDAyNjIwLTIwMTALAAAAAAAAAHRyZWUtci1sYXN0AUmSJAkAAAAAAgAAAAAAAAAhAAAAAAAA
AC9kYXRhLzEvc3RvcmUvY2FjaGUvcy10MDI2MjAtMjAxMAYAAAAAAAAAdHJlZS1jAUmSJAkAAAAA
AAAAAAAAAAA=`

func TestTAuxRoundtrip(t *testing.T) {
	gbytes := must.One(base64.StdEncoding.DecodeString(TAuxGolden))

	dec, err := DecodeTAux(bytes.NewReader(gbytes))
	require.NoError(t, err)

	// print as json
	enc, err := json.MarshalIndent(dec, "", "  ")
	require.NoError(t, err)
	t.Log(string(enc))

	var reenc bytes.Buffer
	require.NoError(t, EncodeTAux(&reenc, *dec))

	require.Equal(t, gbytes, reenc.Bytes())
}