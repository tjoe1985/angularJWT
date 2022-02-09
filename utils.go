package main

//
//func populateWoInDB() {
//	var wo WO
//
//	for i := 10; i < 58; i++ {
//		wo = WO{
//			WoUuid:                  "cda6498a235d4f7eae19661d41bc1" + strconv.Itoa(i) + "c",
//			Client:                  "cda6498a235d4f7eae19661d41bc154d",
//			WorkClient:              "cda6498a235d4f7eae19661d41bc154e",
//			WoNumber:                "OB0000000000" + strconv.Itoa(i),
//			WoRecipient:             "cda6498a235d4f7eae19661d41bc154f",
//			WoNeedsClientAcceptance: false,
//			WoClientAccepted:        false,
//		}
//		dbInsert(ConnectDB(), wo)
//		time.Sleep(2 * time.Second)
//		log.Println("============>>>>>>>>>>>> ", i)
//	}
//
//}
