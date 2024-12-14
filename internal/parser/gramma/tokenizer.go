
//line tokenizer.go.rl:1
// Generated code. DO NOT EDIT

package gramma

import (
	"fmt"
	"strconv"
	"time"
	"errors"
)


//line tokenizer.go:14
const logparser_start int = 1
const logparser_first_final int = 678
const logparser_error int = 0

const logparser_en_main int = 1


//line tokenizer.go.rl:20


type state struct {
	data string

	prev int

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process.
	// It should be initialized to the beginning of the data block on every run of the machine.
	// In Java and Ruby it is used as an offset to data and must be an integer.
	// In this case it should be initialized to zero on every run of the machine.
	p int

	// Token start
	ts int
	// Token end
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state.
	// This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently.
	// This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int
}

func (e state) String() string {
	return fmt.Sprintf("state: %q p=%d, pe=%d, ts=%d, te=%d, act=%d, cs=%d, top=%d", e.data[:e.p], e.p, e.pe, e.ts, e.te, e.act, e.cs, e.top)
}

func (e state) Error() string {
	return e.String()
}

type tokenizer struct {
	nowTime time.Time
	tokens []Token
	errors []error
	state
}


//line tokenizer.go.rl:210


func (t *tokenizer) Parse(nowTime time.Time, data string) ([]Token, error) {
	var parseErr error
	var temp yySymType

	t.state = state{}
	
//line tokenizer.go:78
	{
	 t.cs = logparser_start
	}

//line tokenizer.go.rl:218

	t.prev = 0
	t.pe = len(data)
	t.nowTime = nowTime
	t.data = data

	t.tokens = t.tokens[:0]
	t.errors = nil

	
//line tokenizer.go:92
	{
	if ( t.p) == ( t.pe) {
		goto _test_eof
	}
	switch  t.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 678:
		goto st_case_678
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 679:
		goto st_case_679
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 680:
		goto st_case_680
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 681:
		goto st_case_681
	case 682:
		goto st_case_682
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 683:
		goto st_case_683
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 684:
		goto st_case_684
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 685:
		goto st_case_685
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 686:
		goto st_case_686
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 687:
		goto st_case_687
	case 688:
		goto st_case_688
	case 689:
		goto st_case_689
	case 690:
		goto st_case_690
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 691:
		goto st_case_691
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 692:
		goto st_case_692
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 693:
		goto st_case_693
	case 694:
		goto st_case_694
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 697:
		goto st_case_697
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 652:
		goto st_case_652
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 656:
		goto st_case_656
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 674:
		goto st_case_674
	case 675:
		goto st_case_675
	case 676:
		goto st_case_676
	case 677:
		goto st_case_677
	}
	goto st_out
	st_case_1:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr0
		}
		goto st0
st_case_0:
	st0:
		 t.cs = 0
		goto _out
tr0:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st2
	st2:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof2
		}
	st_case_2:
//line tokenizer.go:1519
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st3
		}
		goto st0
	st3:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof3
		}
	st_case_3:
		if ( t.data)[( t.p)] == 58 {
			goto st4
		}
		goto st0
	st4:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st6
		}
		goto st0
	st6:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof6
		}
	st_case_6:
		if ( t.data)[( t.p)] == 58 {
			goto st7
		}
		goto st0
	st7:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof9
		}
	st_case_9:
		if ( t.data)[( t.p)] == 46 {
			goto st10
		}
		goto st0
	st10:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof13
		}
	st_case_13:
		if ( t.data)[( t.p)] == 32 {
			goto tr13
		}
		goto st0
tr13:
//line tokenizer.go.rl:95

		temp.Time, parseErr = t.parseTime(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: TIME, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("parseTime: %w", parseErr)})
			{( t.p)++;  t.cs = 14; goto _out }
		}
		t.tokval(TimeTok(temp.Time))
	
	goto st14
	st14:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof14
		}
	st_case_14:
//line tokenizer.go:1639
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr15
		case 67:
			goto tr16
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr14
		}
		goto st0
tr14:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st15
	st15:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof15
		}
	st_case_15:
//line tokenizer.go:1664
		switch ( t.data)[( t.p)] {
		case 32:
			goto st15
		case 124:
			goto tr18
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st15
		}
		goto st0
tr18:
//line tokenizer.go.rl:118
 t.tok(GAME) 
	goto st16
	st16:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof16
		}
	st_case_16:
//line tokenizer.go:1684
		if ( t.data)[( t.p)] == 32 {
			goto st17
		}
		goto st0
	st17:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st17
		case 99:
			goto tr20
		}
		goto st0
tr20:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st18
	st18:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof18
		}
	st_case_18:
//line tokenizer.go:1715
		if ( t.data)[( t.p)] == 108 {
			goto st19
		}
		goto st0
	st19:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof19
		}
	st_case_19:
		if ( t.data)[( t.p)] == 105 {
			goto st20
		}
		goto st0
	st20:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof20
		}
	st_case_20:
		if ( t.data)[( t.p)] == 101 {
			goto st21
		}
		goto st0
	st21:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof21
		}
	st_case_21:
		if ( t.data)[( t.p)] == 110 {
			goto st22
		}
		goto st0
	st22:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof22
		}
	st_case_22:
		if ( t.data)[( t.p)] == 116 {
			goto st23
		}
		goto st0
	st23:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof23
		}
	st_case_23:
		if ( t.data)[( t.p)] == 58 {
			goto st24
		}
		goto st0
	st24:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof24
		}
	st_case_24:
		if ( t.data)[( t.p)] == 32 {
			goto st25
		}
		goto st0
	st25:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch ( t.data)[( t.p)] {
		case 65:
			goto st26
		case 99:
			goto st81
		case 112:
			goto st155
		}
		goto st0
	st26:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof26
		}
	st_case_26:
		if ( t.data)[( t.p)] == 68 {
			goto st27
		}
		goto st0
	st27:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof27
		}
	st_case_27:
		if ( t.data)[( t.p)] == 68 {
			goto st28
		}
		goto st0
	st28:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof28
		}
	st_case_28:
		if ( t.data)[( t.p)] == 95 {
			goto st29
		}
		goto st0
	st29:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof29
		}
	st_case_29:
		if ( t.data)[( t.p)] == 80 {
			goto st30
		}
		goto st0
	st30:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof30
		}
	st_case_30:
		if ( t.data)[( t.p)] == 76 {
			goto st31
		}
		goto st0
	st31:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof31
		}
	st_case_31:
		if ( t.data)[( t.p)] == 65 {
			goto st32
		}
		goto st0
	st32:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof32
		}
	st_case_32:
		if ( t.data)[( t.p)] == 89 {
			goto st33
		}
		goto st0
	st33:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof33
		}
	st_case_33:
		if ( t.data)[( t.p)] == 69 {
			goto st34
		}
		goto st0
	st34:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof34
		}
	st_case_34:
		if ( t.data)[( t.p)] == 82 {
			goto st35
		}
		goto st0
	st35:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof35
		}
	st_case_35:
		if ( t.data)[( t.p)] == 32 {
			goto st36
		}
		goto st0
	st36:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof36
		}
	st_case_36:
		if ( t.data)[( t.p)] == 45 {
			goto tr41
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr42
		}
		goto st0
tr41:
//line tokenizer.go.rl:178
t.tok(CLIENT_ADD_PLAYER)
	goto st37
	st37:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof37
		}
	st_case_37:
//line tokenizer.go:1899
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr43
		}
		goto st0
tr43:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st38
tr42:
//line tokenizer.go.rl:178
t.tok(CLIENT_ADD_PLAYER)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st38
	st38:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof38
		}
	st_case_38:
//line tokenizer.go:1929
		if ( t.data)[( t.p)] == 40 {
			goto tr44
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st38
		}
		goto st0
tr44:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 39; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st39
	st39:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof39
		}
	st_case_39:
//line tokenizer.go:1953
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr46
		case 95:
			goto tr47
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr47
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr47
			}
		default:
			goto tr47
		}
		goto st0
tr46:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st40
	st40:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof40
		}
	st_case_40:
//line tokenizer.go:1987
		if ( t.data)[( t.p)] == 95 {
			goto st41
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st41
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto st0
	st41:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st41
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto st0
	st42:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof42
		}
	st_case_42:
		if ( t.data)[( t.p)] == 32 {
			goto tr50
		}
		goto st0
tr50:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st43
	st43:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof43
		}
	st_case_43:
//line tokenizer.go:2048
		if ( t.data)[( t.p)] == 91 {
			goto st44
		}
		goto st0
	st44:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr52
		case 95:
			goto tr53
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr53
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr53
			}
		default:
			goto tr53
		}
		goto st0
tr52:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st45
	st45:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof45
		}
	st_case_45:
//line tokenizer.go:2091
		if ( t.data)[( t.p)] == 95 {
			goto st46
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st46
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st46
			}
		default:
			goto st46
		}
		goto st0
	st46:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st47
		case 95:
			goto st46
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st46
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st46
			}
		default:
			goto st46
		}
		goto st0
	st47:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof47
		}
	st_case_47:
		if ( t.data)[( t.p)] == 93 {
			goto tr56
		}
		goto st0
tr56:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st48
	st48:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof48
		}
	st_case_48:
//line tokenizer.go:2152
		if ( t.data)[( t.p)] == 44 {
			goto st49
		}
		goto st0
	st49:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof49
		}
	st_case_49:
		if ( t.data)[( t.p)] == 32 {
			goto st50
		}
		goto st0
	st50:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof50
		}
	st_case_50:
		if ( t.data)[( t.p)] == 45 {
			goto st51
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr60
		}
		goto st0
	st51:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof51
		}
	st_case_51:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr60
		}
		goto st0
tr60:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st52
	st52:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof52
		}
	st_case_52:
//line tokenizer.go:2201
		if ( t.data)[( t.p)] == 41 {
			goto tr61
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st52
		}
		goto st0
tr61:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 53; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st53
	st53:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof53
		}
	st_case_53:
//line tokenizer.go:2225
		if ( t.data)[( t.p)] == 32 {
			goto st54
		}
		goto st0
	st54:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof54
		}
	st_case_54:
		if ( t.data)[( t.p)] == 115 {
			goto st55
		}
		goto st0
	st55:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof55
		}
	st_case_55:
		if ( t.data)[( t.p)] == 116 {
			goto st56
		}
		goto st0
	st56:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof56
		}
	st_case_56:
		if ( t.data)[( t.p)] == 97 {
			goto st57
		}
		goto st0
	st57:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof57
		}
	st_case_57:
		if ( t.data)[( t.p)] == 116 {
			goto st58
		}
		goto st0
	st58:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof58
		}
	st_case_58:
		if ( t.data)[( t.p)] == 117 {
			goto st59
		}
		goto st0
	st59:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof59
		}
	st_case_59:
		if ( t.data)[( t.p)] == 115 {
			goto st60
		}
		goto st0
	st60:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof60
		}
	st_case_60:
		if ( t.data)[( t.p)] == 32 {
			goto st61
		}
		goto st0
	st61:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof61
		}
	st_case_61:
		if ( t.data)[( t.p)] == 45 {
			goto st62
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr72
		}
		goto st0
	st62:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof62
		}
	st_case_62:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr72
		}
		goto st0
tr72:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st63
	st63:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof63
		}
	st_case_63:
//line tokenizer.go:2328
		if ( t.data)[( t.p)] == 32 {
			goto tr73
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st63
		}
		goto st0
tr73:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 64; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st64
	st64:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof64
		}
	st_case_64:
//line tokenizer.go:2352
		if ( t.data)[( t.p)] == 116 {
			goto st65
		}
		goto st0
	st65:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof65
		}
	st_case_65:
		if ( t.data)[( t.p)] == 101 {
			goto st66
		}
		goto st0
	st66:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof66
		}
	st_case_66:
		if ( t.data)[( t.p)] == 97 {
			goto st67
		}
		goto st0
	st67:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof67
		}
	st_case_67:
		if ( t.data)[( t.p)] == 109 {
			goto st68
		}
		goto st0
	st68:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof68
		}
	st_case_68:
		if ( t.data)[( t.p)] == 32 {
			goto st69
		}
		goto st0
	st69:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof69
		}
	st_case_69:
		if ( t.data)[( t.p)] == 45 {
			goto st70
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr81
		}
		goto st0
	st70:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof70
		}
	st_case_70:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr81
		}
		goto st0
tr81:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st678
	st678:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof678
		}
	st_case_678:
//line tokenizer.go:2428
		if ( t.data)[( t.p)] == 32 {
			goto tr756
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st678
		}
		goto st0
tr756:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 71; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st71
	st71:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof71
		}
	st_case_71:
//line tokenizer.go:2452
		if ( t.data)[( t.p)] == 103 {
			goto st72
		}
		goto st0
	st72:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof72
		}
	st_case_72:
		if ( t.data)[( t.p)] == 114 {
			goto st73
		}
		goto st0
	st73:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof73
		}
	st_case_73:
		if ( t.data)[( t.p)] == 111 {
			goto st74
		}
		goto st0
	st74:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof74
		}
	st_case_74:
		if ( t.data)[( t.p)] == 117 {
			goto st75
		}
		goto st0
	st75:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof75
		}
	st_case_75:
		if ( t.data)[( t.p)] == 112 {
			goto st76
		}
		goto st0
	st76:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof76
		}
	st_case_76:
		if ( t.data)[( t.p)] == 32 {
			goto st77
		}
		goto st0
	st77:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof77
		}
	st_case_77:
		if ( t.data)[( t.p)] == 45 {
			goto st78
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr89
		}
		goto st0
	st78:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof78
		}
	st_case_78:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr89
		}
		goto st0
tr89:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st679
	st679:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof679
		}
	st_case_679:
//line tokenizer.go:2537
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st679
		}
		goto st0
tr53:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st79
	st79:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof79
		}
	st_case_79:
//line tokenizer.go:2556
		switch ( t.data)[( t.p)] {
		case 93:
			goto tr56
		case 95:
			goto st79
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st79
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st79
			}
		default:
			goto st79
		}
		goto st0
tr47:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st80
	st80:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof80
		}
	st_case_80:
//line tokenizer.go:2590
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr50
		case 95:
			goto st80
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st80
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st81:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof81
		}
	st_case_81:
		if ( t.data)[( t.p)] == 111 {
			goto st82
		}
		goto st0
	st82:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof82
		}
	st_case_82:
		if ( t.data)[( t.p)] == 110 {
			goto st83
		}
		goto st0
	st83:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof83
		}
	st_case_83:
		if ( t.data)[( t.p)] == 110 {
			goto st84
		}
		goto st0
	st84:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof84
		}
	st_case_84:
		if ( t.data)[( t.p)] == 101 {
			goto st85
		}
		goto st0
	st85:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof85
		}
	st_case_85:
		if ( t.data)[( t.p)] == 99 {
			goto st86
		}
		goto st0
	st86:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof86
		}
	st_case_86:
		if ( t.data)[( t.p)] == 116 {
			goto st87
		}
		goto st0
	st87:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch ( t.data)[( t.p)] {
		case 101:
			goto st88
		case 105:
			goto st141
		}
		goto st0
	st88:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof88
		}
	st_case_88:
		if ( t.data)[( t.p)] == 100 {
			goto st89
		}
		goto st0
	st89:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof89
		}
	st_case_89:
		if ( t.data)[( t.p)] == 32 {
			goto st90
		}
		goto st0
	st90:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof90
		}
	st_case_90:
		if ( t.data)[( t.p)] == 116 {
			goto st91
		}
		goto st0
	st91:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof91
		}
	st_case_91:
		if ( t.data)[( t.p)] == 111 {
			goto st92
		}
		goto st0
	st92:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof92
		}
	st_case_92:
		if ( t.data)[( t.p)] == 32 {
			goto st93
		}
		goto st0
	st93:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof93
		}
	st_case_93:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr105
		}
		goto st0
tr105:
//line tokenizer.go.rl:181
t.tok(CLIENT_CONNECTED)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st94
	st94:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof94
		}
	st_case_94:
//line tokenizer.go:2746
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st139
		}
		goto st0
	st95:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof95
		}
	st_case_95:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st96
		}
		goto st0
	st96:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof96
		}
	st_case_96:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st137
		}
		goto st0
	st97:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof97
		}
	st_case_97:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st98
		}
		goto st0
	st98:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof98
		}
	st_case_98:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st135
		}
		goto st0
	st99:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof99
		}
	st_case_99:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st100
		}
		goto st0
	st100:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof100
		}
	st_case_100:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st101
		}
		goto st0
	st101:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof101
		}
	st_case_101:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st102
		}
		goto st0
	st102:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof102
		}
	st_case_102:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		goto st0
	st103:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof103
		}
	st_case_103:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st104
		}
		goto st0
	st104:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof104
		}
	st_case_104:
		if ( t.data)[( t.p)] == 44 {
			goto tr119
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st104
		}
		goto st0
tr119:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st105
	st105:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof105
		}
	st_case_105:
//line tokenizer.go:2870
		if ( t.data)[( t.p)] == 32 {
			goto st106
		}
		goto st0
	st106:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof106
		}
	st_case_106:
		if ( t.data)[( t.p)] == 77 {
			goto st107
		}
		goto st0
	st107:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof107
		}
	st_case_107:
		if ( t.data)[( t.p)] == 84 {
			goto st108
		}
		goto st0
	st108:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof108
		}
	st_case_108:
		if ( t.data)[( t.p)] == 85 {
			goto st109
		}
		goto st0
	st109:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof109
		}
	st_case_109:
		if ( t.data)[( t.p)] == 32 {
			goto st110
		}
		goto st0
	st110:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof110
		}
	st_case_110:
		if ( t.data)[( t.p)] == 45 {
			goto st111
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr126
		}
		goto st0
	st111:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof111
		}
	st_case_111:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr126
		}
		goto st0
tr126:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st112
	st112:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof112
		}
	st_case_112:
//line tokenizer.go:2946
		if ( t.data)[( t.p)] == 46 {
			goto tr127
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st112
		}
		goto st0
tr127:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 113; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st113
	st113:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof113
		}
	st_case_113:
//line tokenizer.go:2970
		if ( t.data)[( t.p)] == 32 {
			goto st114
		}
		goto st0
	st114:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof114
		}
	st_case_114:
		if ( t.data)[( t.p)] == 115 {
			goto st115
		}
		goto st0
	st115:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof115
		}
	st_case_115:
		if ( t.data)[( t.p)] == 101 {
			goto st116
		}
		goto st0
	st116:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof116
		}
	st_case_116:
		if ( t.data)[( t.p)] == 116 {
			goto st117
		}
		goto st0
	st117:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof117
		}
	st_case_117:
		if ( t.data)[( t.p)] == 116 {
			goto st118
		}
		goto st0
	st118:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof118
		}
	st_case_118:
		if ( t.data)[( t.p)] == 105 {
			goto st119
		}
		goto st0
	st119:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof119
		}
	st_case_119:
		if ( t.data)[( t.p)] == 110 {
			goto st120
		}
		goto st0
	st120:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof120
		}
	st_case_120:
		if ( t.data)[( t.p)] == 103 {
			goto st121
		}
		goto st0
	st121:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof121
		}
	st_case_121:
		if ( t.data)[( t.p)] == 32 {
			goto st122
		}
		goto st0
	st122:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof122
		}
	st_case_122:
		if ( t.data)[( t.p)] == 117 {
			goto st123
		}
		goto st0
	st123:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof123
		}
	st_case_123:
		if ( t.data)[( t.p)] == 112 {
			goto st124
		}
		goto st0
	st124:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof124
		}
	st_case_124:
		if ( t.data)[( t.p)] == 32 {
			goto st125
		}
		goto st0
	st125:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof125
		}
	st_case_125:
		if ( t.data)[( t.p)] == 115 {
			goto st126
		}
		goto st0
	st126:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof126
		}
	st_case_126:
		if ( t.data)[( t.p)] == 101 {
			goto st127
		}
		goto st0
	st127:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof127
		}
	st_case_127:
		if ( t.data)[( t.p)] == 115 {
			goto st128
		}
		goto st0
	st128:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof128
		}
	st_case_128:
		if ( t.data)[( t.p)] == 115 {
			goto st129
		}
		goto st0
	st129:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof129
		}
	st_case_129:
		if ( t.data)[( t.p)] == 105 {
			goto st130
		}
		goto st0
	st130:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof130
		}
	st_case_130:
		if ( t.data)[( t.p)] == 111 {
			goto st131
		}
		goto st0
	st131:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof131
		}
	st_case_131:
		if ( t.data)[( t.p)] == 110 {
			goto st132
		}
		goto st0
	st132:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof132
		}
	st_case_132:
		if ( t.data)[( t.p)] == 46 {
			goto st133
		}
		goto st0
	st133:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof133
		}
	st_case_133:
		if ( t.data)[( t.p)] == 46 {
			goto st134
		}
		goto st0
	st134:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof134
		}
	st_case_134:
		if ( t.data)[( t.p)] == 46 {
			goto st680
		}
		goto st0
tr548:
//line tokenizer.go.rl:138
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st680
tr610:
//line tokenizer.go.rl:151
t.tok(FRIENDLY_FIRE)
	goto st680
	st680:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof680
		}
	st_case_680:
//line tokenizer.go:3177
		goto st0
	st135:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof135
		}
	st_case_135:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st136
		}
		goto st0
	st136:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof136
		}
	st_case_136:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		goto st0
	st137:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof137
		}
	st_case_137:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st138
		}
		goto st0
	st138:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof138
		}
	st_case_138:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		goto st0
	st139:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof139
		}
	st_case_139:
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st140
		}
		goto st0
	st140:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof140
		}
	st_case_140:
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		goto st0
	st141:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof141
		}
	st_case_141:
		if ( t.data)[( t.p)] == 111 {
			goto st142
		}
		goto st0
	st142:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof142
		}
	st_case_142:
		if ( t.data)[( t.p)] == 110 {
			goto st143
		}
		goto st0
	st143:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof143
		}
	st_case_143:
		if ( t.data)[( t.p)] == 32 {
			goto st144
		}
		goto st0
	st144:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof144
		}
	st_case_144:
		if ( t.data)[( t.p)] == 99 {
			goto st145
		}
		goto st0
	st145:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof145
		}
	st_case_145:
		if ( t.data)[( t.p)] == 108 {
			goto st146
		}
		goto st0
	st146:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof146
		}
	st_case_146:
		if ( t.data)[( t.p)] == 111 {
			goto st147
		}
		goto st0
	st147:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof147
		}
	st_case_147:
		if ( t.data)[( t.p)] == 115 {
			goto st148
		}
		goto st0
	st148:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof148
		}
	st_case_148:
		if ( t.data)[( t.p)] == 101 {
			goto st149
		}
		goto st0
	st149:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof149
		}
	st_case_149:
		if ( t.data)[( t.p)] == 100 {
			goto st150
		}
		goto st0
	st150:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof150
		}
	st_case_150:
		if ( t.data)[( t.p)] == 46 {
			goto st151
		}
		goto st0
	st151:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof151
		}
	st_case_151:
		if ( t.data)[( t.p)] == 32 {
			goto st152
		}
		goto st0
	st152:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof152
		}
	st_case_152:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr165
		case 95:
			goto tr166
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr166
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr166
			}
		default:
			goto tr166
		}
		goto st0
tr165:
//line tokenizer.go.rl:184
t.tok(CLIENT_CONNECTION_CLOSED)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st153
	st153:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof153
		}
	st_case_153:
//line tokenizer.go:3381
		if ( t.data)[( t.p)] == 95 {
			goto st154
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st154
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto st0
	st154:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st681
		case 95:
			goto st154
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st154
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto st0
tr724:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st681
	st681:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof681
		}
	st_case_681:
//line tokenizer.go:3436
		goto st0
tr166:
//line tokenizer.go.rl:184
t.tok(CLIENT_CONNECTION_CLOSED)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st682
	st682:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof682
		}
	st_case_682:
//line tokenizer.go:3454
		if ( t.data)[( t.p)] == 95 {
			goto st682
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st682
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st682
			}
		default:
			goto st682
		}
		goto st0
	st155:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof155
		}
	st_case_155:
		if ( t.data)[( t.p)] == 108 {
			goto st156
		}
		goto st0
	st156:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof156
		}
	st_case_156:
		if ( t.data)[( t.p)] == 97 {
			goto st157
		}
		goto st0
	st157:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof157
		}
	st_case_157:
		if ( t.data)[( t.p)] == 121 {
			goto st158
		}
		goto st0
	st158:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof158
		}
	st_case_158:
		if ( t.data)[( t.p)] == 101 {
			goto st159
		}
		goto st0
	st159:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof159
		}
	st_case_159:
		if ( t.data)[( t.p)] == 114 {
			goto st160
		}
		goto st0
	st160:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof160
		}
	st_case_160:
		if ( t.data)[( t.p)] == 32 {
			goto st161
		}
		goto st0
	st161:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof161
		}
	st_case_161:
		if ( t.data)[( t.p)] == 45 {
			goto tr175
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr176
		}
		goto st0
tr175:
//line tokenizer.go.rl:180
t.tok(CLIENT_PLAYER_LEAVE)
	goto st162
	st162:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof162
		}
	st_case_162:
//line tokenizer.go:3546
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr177
		}
		goto st0
tr177:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st163
tr176:
//line tokenizer.go.rl:180
t.tok(CLIENT_PLAYER_LEAVE)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st163
	st163:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof163
		}
	st_case_163:
//line tokenizer.go:3576
		if ( t.data)[( t.p)] == 32 {
			goto tr178
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st163
		}
		goto st0
tr178:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 164; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st164
	st164:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof164
		}
	st_case_164:
//line tokenizer.go:3600
		if ( t.data)[( t.p)] == 108 {
			goto st165
		}
		goto st0
	st165:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof165
		}
	st_case_165:
		if ( t.data)[( t.p)] == 101 {
			goto st166
		}
		goto st0
	st166:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof166
		}
	st_case_166:
		if ( t.data)[( t.p)] == 97 {
			goto st167
		}
		goto st0
	st167:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof167
		}
	st_case_167:
		if ( t.data)[( t.p)] == 118 {
			goto st168
		}
		goto st0
	st168:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof168
		}
	st_case_168:
		if ( t.data)[( t.p)] == 101 {
			goto st169
		}
		goto st0
	st169:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof169
		}
	st_case_169:
		if ( t.data)[( t.p)] == 32 {
			goto st170
		}
		goto st0
	st170:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof170
		}
	st_case_170:
		if ( t.data)[( t.p)] == 103 {
			goto st171
		}
		goto st0
	st171:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof171
		}
	st_case_171:
		if ( t.data)[( t.p)] == 97 {
			goto st172
		}
		goto st0
	st172:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof172
		}
	st_case_172:
		if ( t.data)[( t.p)] == 109 {
			goto st173
		}
		goto st0
	st173:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof173
		}
	st_case_173:
		if ( t.data)[( t.p)] == 101 {
			goto st680
		}
		goto st0
tr15:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st174
	st174:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof174
		}
	st_case_174:
//line tokenizer.go:3700
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr15
		case 67:
			goto tr16
		case 124:
			goto tr18
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr14
		}
		goto st0
tr16:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st175
	st175:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof175
		}
	st_case_175:
//line tokenizer.go:3727
		if ( t.data)[( t.p)] == 77 {
			goto st176
		}
		goto st0
	st176:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof176
		}
	st_case_176:
		if ( t.data)[( t.p)] == 66 {
			goto st177
		}
		goto st0
	st177:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof177
		}
	st_case_177:
		if ( t.data)[( t.p)] == 84 {
			goto st178
		}
		goto st0
	st178:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof178
		}
	st_case_178:
		if ( t.data)[( t.p)] == 32 {
			goto st179
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st179
		}
		goto st0
	st179:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st179
		case 124:
			goto tr193
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st179
		}
		goto st0
tr193:
//line tokenizer.go.rl:117
 t.tok(COMBAT) 
	goto st180
	st180:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof180
		}
	st_case_180:
//line tokenizer.go:3786
		if ( t.data)[( t.p)] == 32 {
			goto st181
		}
		goto st0
	st181:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st181
		case 61:
			goto tr195
		case 68:
			goto tr196
		case 71:
			goto tr197
		case 72:
			goto tr198
		case 75:
			goto tr199
		case 80:
			goto tr200
		case 82:
			goto tr201
		}
		goto st0
tr195:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st182
	st182:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof182
		}
	st_case_182:
//line tokenizer.go:3829
		if ( t.data)[( t.p)] == 61 {
			goto st183
		}
		goto st0
	st183:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof183
		}
	st_case_183:
		if ( t.data)[( t.p)] == 61 {
			goto st184
		}
		goto st0
	st184:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof184
		}
	st_case_184:
		if ( t.data)[( t.p)] == 61 {
			goto st185
		}
		goto st0
	st185:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof185
		}
	st_case_185:
		if ( t.data)[( t.p)] == 61 {
			goto st186
		}
		goto st0
	st186:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof186
		}
	st_case_186:
		if ( t.data)[( t.p)] == 61 {
			goto st187
		}
		goto st0
	st187:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof187
		}
	st_case_187:
		if ( t.data)[( t.p)] == 61 {
			goto st188
		}
		goto st0
	st188:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof188
		}
	st_case_188:
		if ( t.data)[( t.p)] == 32 {
			goto st189
		}
		goto st0
	st189:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch ( t.data)[( t.p)] {
		case 67:
			goto st190
		case 83:
			goto st223
		}
		goto st0
	st190:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof190
		}
	st_case_190:
		if ( t.data)[( t.p)] == 111 {
			goto st191
		}
		goto st0
	st191:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof191
		}
	st_case_191:
		if ( t.data)[( t.p)] == 110 {
			goto st192
		}
		goto st0
	st192:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof192
		}
	st_case_192:
		if ( t.data)[( t.p)] == 110 {
			goto st193
		}
		goto st0
	st193:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof193
		}
	st_case_193:
		if ( t.data)[( t.p)] == 101 {
			goto st194
		}
		goto st0
	st194:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof194
		}
	st_case_194:
		if ( t.data)[( t.p)] == 99 {
			goto st195
		}
		goto st0
	st195:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof195
		}
	st_case_195:
		if ( t.data)[( t.p)] == 116 {
			goto st196
		}
		goto st0
	st196:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof196
		}
	st_case_196:
		if ( t.data)[( t.p)] == 32 {
			goto st197
		}
		goto st0
	st197:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof197
		}
	st_case_197:
		if ( t.data)[( t.p)] == 116 {
			goto st198
		}
		goto st0
	st198:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof198
		}
	st_case_198:
		if ( t.data)[( t.p)] == 111 {
			goto st199
		}
		goto st0
	st199:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof199
		}
	st_case_199:
		if ( t.data)[( t.p)] == 32 {
			goto st200
		}
		goto st0
	st200:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof200
		}
	st_case_200:
		if ( t.data)[( t.p)] == 103 {
			goto st201
		}
		goto st0
	st201:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof201
		}
	st_case_201:
		if ( t.data)[( t.p)] == 97 {
			goto st202
		}
		goto st0
	st202:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof202
		}
	st_case_202:
		if ( t.data)[( t.p)] == 109 {
			goto st203
		}
		goto st0
	st203:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof203
		}
	st_case_203:
		if ( t.data)[( t.p)] == 101 {
			goto st204
		}
		goto st0
	st204:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof204
		}
	st_case_204:
		if ( t.data)[( t.p)] == 32 {
			goto st205
		}
		goto st0
	st205:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof205
		}
	st_case_205:
		if ( t.data)[( t.p)] == 115 {
			goto st206
		}
		goto st0
	st206:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof206
		}
	st_case_206:
		if ( t.data)[( t.p)] == 101 {
			goto st207
		}
		goto st0
	st207:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof207
		}
	st_case_207:
		if ( t.data)[( t.p)] == 115 {
			goto st208
		}
		goto st0
	st208:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof208
		}
	st_case_208:
		if ( t.data)[( t.p)] == 115 {
			goto st209
		}
		goto st0
	st209:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof209
		}
	st_case_209:
		if ( t.data)[( t.p)] == 105 {
			goto st210
		}
		goto st0
	st210:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof210
		}
	st_case_210:
		if ( t.data)[( t.p)] == 111 {
			goto st211
		}
		goto st0
	st211:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof211
		}
	st_case_211:
		if ( t.data)[( t.p)] == 110 {
			goto st212
		}
		goto st0
	st212:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof212
		}
	st_case_212:
		if ( t.data)[( t.p)] == 32 {
			goto st213
		}
		goto st0
	st213:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof213
		}
	st_case_213:
		if ( t.data)[( t.p)] == 45 {
			goto tr234
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr235
		}
		goto st0
tr234:
//line tokenizer.go.rl:173
t.tok(CONNECT_TO_GAME_SESSION_PREFIX)
	goto st214
	st214:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof214
		}
	st_case_214:
//line tokenizer.go:4128
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr236
		}
		goto st0
tr236:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st215
tr235:
//line tokenizer.go.rl:173
t.tok(CONNECT_TO_GAME_SESSION_PREFIX)
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st215
	st215:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof215
		}
	st_case_215:
//line tokenizer.go:4158
		if ( t.data)[( t.p)] == 32 {
			goto tr237
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st215
		}
		goto st0
tr237:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 216; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st216
	st216:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof216
		}
	st_case_216:
//line tokenizer.go:4182
		if ( t.data)[( t.p)] == 61 {
			goto st217
		}
		goto st0
tr301:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 217; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st217
	st217:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof217
		}
	st_case_217:
//line tokenizer.go:4203
		if ( t.data)[( t.p)] == 61 {
			goto st218
		}
		goto st0
	st218:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof218
		}
	st_case_218:
		if ( t.data)[( t.p)] == 61 {
			goto st219
		}
		goto st0
	st219:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof219
		}
	st_case_219:
		if ( t.data)[( t.p)] == 61 {
			goto st220
		}
		goto st0
	st220:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof220
		}
	st_case_220:
		if ( t.data)[( t.p)] == 61 {
			goto st221
		}
		goto st0
	st221:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof221
		}
	st_case_221:
		if ( t.data)[( t.p)] == 61 {
			goto st222
		}
		goto st0
	st222:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof222
		}
	st_case_222:
		if ( t.data)[( t.p)] == 61 {
			goto st680
		}
		goto st0
	st223:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof223
		}
	st_case_223:
		if ( t.data)[( t.p)] == 116 {
			goto st224
		}
		goto st0
	st224:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof224
		}
	st_case_224:
		if ( t.data)[( t.p)] == 97 {
			goto st225
		}
		goto st0
	st225:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof225
		}
	st_case_225:
		if ( t.data)[( t.p)] == 114 {
			goto st226
		}
		goto st0
	st226:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof226
		}
	st_case_226:
		if ( t.data)[( t.p)] == 116 {
			goto tr248
		}
		goto st0
tr248:
//line tokenizer.go.rl:125
 t.tok(START) 
	goto st227
	st227:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof227
		}
	st_case_227:
//line tokenizer.go:4298
		if ( t.data)[( t.p)] == 32 {
			goto st228
		}
		goto st0
	st228:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof228
		}
	st_case_228:
		switch ( t.data)[( t.p)] {
		case 80:
			goto tr250
		case 103:
			goto tr251
		}
		goto st0
tr250:
//line tokenizer.go.rl:169
t.tokval(StrTok("PVE mission"))
	goto st229
	st229:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof229
		}
	st_case_229:
//line tokenizer.go:4324
		if ( t.data)[( t.p)] == 86 {
			goto st230
		}
		goto st0
	st230:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof230
		}
	st_case_230:
		if ( t.data)[( t.p)] == 69 {
			goto st231
		}
		goto st0
	st231:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof231
		}
	st_case_231:
		if ( t.data)[( t.p)] == 32 {
			goto st232
		}
		goto st0
	st232:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof232
		}
	st_case_232:
		if ( t.data)[( t.p)] == 109 {
			goto st233
		}
		goto st0
	st233:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof233
		}
	st_case_233:
		if ( t.data)[( t.p)] == 105 {
			goto st234
		}
		goto st0
	st234:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof234
		}
	st_case_234:
		if ( t.data)[( t.p)] == 115 {
			goto st235
		}
		goto st0
	st235:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof235
		}
	st_case_235:
		if ( t.data)[( t.p)] == 115 {
			goto st236
		}
		goto st0
	st236:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof236
		}
	st_case_236:
		if ( t.data)[( t.p)] == 105 {
			goto st237
		}
		goto st0
	st237:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof237
		}
	st_case_237:
		if ( t.data)[( t.p)] == 111 {
			goto st238
		}
		goto st0
	st238:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof238
		}
	st_case_238:
		if ( t.data)[( t.p)] == 110 {
			goto st239
		}
		goto st0
	st239:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof239
		}
	st_case_239:
		if ( t.data)[( t.p)] == 32 {
			goto st240
		}
		goto st0
	st240:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof240
		}
	st_case_240:
		if ( t.data)[( t.p)] == 39 {
			goto st241
		}
		goto st0
	st241:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr264
		case 95:
			goto tr265
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr265
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr265
			}
		default:
			goto tr265
		}
		goto st0
tr264:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st242
	st242:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof242
		}
	st_case_242:
//line tokenizer.go:4466
		if ( t.data)[( t.p)] == 95 {
			goto st243
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st243
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st243
			}
		default:
			goto st243
		}
		goto st0
	st243:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st244
		case 95:
			goto st243
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st243
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st243
			}
		default:
			goto st243
		}
		goto st0
	st244:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof244
		}
	st_case_244:
		if ( t.data)[( t.p)] == 39 {
			goto tr268
		}
		goto st0
tr268:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st245
	st245:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof245
		}
	st_case_245:
//line tokenizer.go:4527
		if ( t.data)[( t.p)] == 109 {
			goto st246
		}
		goto st0
	st246:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof246
		}
	st_case_246:
		if ( t.data)[( t.p)] == 97 {
			goto st247
		}
		goto st0
	st247:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof247
		}
	st_case_247:
		if ( t.data)[( t.p)] == 112 {
			goto st248
		}
		goto st0
	st248:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof248
		}
	st_case_248:
		if ( t.data)[( t.p)] == 32 {
			goto st249
		}
		goto st0
	st249:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof249
		}
	st_case_249:
		if ( t.data)[( t.p)] == 39 {
			goto st250
		}
		goto st0
	st250:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr274
		case 95:
			goto tr275
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr275
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr275
			}
		default:
			goto tr275
		}
		goto st0
tr274:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st251
	st251:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof251
		}
	st_case_251:
//line tokenizer.go:4606
		if ( t.data)[( t.p)] == 95 {
			goto st252
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st252
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st252
			}
		default:
			goto st252
		}
		goto st0
	st252:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st253
		case 95:
			goto st252
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st252
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st252
			}
		default:
			goto st252
		}
		goto st0
	st253:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof253
		}
	st_case_253:
		if ( t.data)[( t.p)] == 39 {
			goto tr278
		}
		goto st0
tr278:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st683
	st683:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof683
		}
	st_case_683:
//line tokenizer.go:4667
		if ( t.data)[( t.p)] == 44 {
			goto st254
		}
		goto st0
	st254:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof254
		}
	st_case_254:
		if ( t.data)[( t.p)] == 32 {
			goto st255
		}
		goto st0
	st255:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof255
		}
	st_case_255:
		if ( t.data)[( t.p)] == 108 {
			goto st256
		}
		goto st0
	st256:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof256
		}
	st_case_256:
		if ( t.data)[( t.p)] == 111 {
			goto st257
		}
		goto st0
	st257:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof257
		}
	st_case_257:
		if ( t.data)[( t.p)] == 99 {
			goto st258
		}
		goto st0
	st258:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof258
		}
	st_case_258:
		if ( t.data)[( t.p)] == 97 {
			goto st259
		}
		goto st0
	st259:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof259
		}
	st_case_259:
		if ( t.data)[( t.p)] == 108 {
			goto st260
		}
		goto st0
	st260:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof260
		}
	st_case_260:
		if ( t.data)[( t.p)] == 32 {
			goto st261
		}
		goto st0
	st261:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof261
		}
	st_case_261:
		if ( t.data)[( t.p)] == 99 {
			goto st262
		}
		goto st0
	st262:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof262
		}
	st_case_262:
		if ( t.data)[( t.p)] == 108 {
			goto st263
		}
		goto st0
	st263:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof263
		}
	st_case_263:
		if ( t.data)[( t.p)] == 105 {
			goto st264
		}
		goto st0
	st264:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof264
		}
	st_case_264:
		if ( t.data)[( t.p)] == 101 {
			goto st265
		}
		goto st0
	st265:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof265
		}
	st_case_265:
		if ( t.data)[( t.p)] == 110 {
			goto st266
		}
		goto st0
	st266:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof266
		}
	st_case_266:
		if ( t.data)[( t.p)] == 116 {
			goto st267
		}
		goto st0
	st267:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof267
		}
	st_case_267:
		if ( t.data)[( t.p)] == 32 {
			goto st268
		}
		goto st0
	st268:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof268
		}
	st_case_268:
		if ( t.data)[( t.p)] == 116 {
			goto st269
		}
		goto st0
	st269:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof269
		}
	st_case_269:
		if ( t.data)[( t.p)] == 101 {
			goto st270
		}
		goto st0
	st270:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof270
		}
	st_case_270:
		if ( t.data)[( t.p)] == 97 {
			goto st271
		}
		goto st0
	st271:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof271
		}
	st_case_271:
		if ( t.data)[( t.p)] == 109 {
			goto st272
		}
		goto st0
	st272:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof272
		}
	st_case_272:
		if ( t.data)[( t.p)] == 32 {
			goto st273
		}
		goto st0
	st273:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof273
		}
	st_case_273:
		if ( t.data)[( t.p)] == 45 {
			goto st274
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr299
		}
		goto st0
	st274:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof274
		}
	st_case_274:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr299
		}
		goto st0
tr299:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st275
	st275:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof275
		}
	st_case_275:
//line tokenizer.go:4878
		if ( t.data)[( t.p)] == 61 {
			goto tr301
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st275
		}
		goto st0
tr275:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st276
	st276:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof276
		}
	st_case_276:
//line tokenizer.go:4900
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr278
		case 95:
			goto st276
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st276
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st276
			}
		default:
			goto st276
		}
		goto st0
tr265:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st277
	st277:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof277
		}
	st_case_277:
//line tokenizer.go:4934
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr268
		case 95:
			goto st277
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st277
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st277
			}
		default:
			goto st277
		}
		goto st0
tr251:
//line tokenizer.go.rl:168
t.tokval(StrTok("gameplay"))
	goto st278
	st278:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof278
		}
	st_case_278:
//line tokenizer.go:4963
		if ( t.data)[( t.p)] == 97 {
			goto st279
		}
		goto st0
	st279:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof279
		}
	st_case_279:
		if ( t.data)[( t.p)] == 109 {
			goto st280
		}
		goto st0
	st280:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof280
		}
	st_case_280:
		if ( t.data)[( t.p)] == 101 {
			goto st281
		}
		goto st0
	st281:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof281
		}
	st_case_281:
		if ( t.data)[( t.p)] == 112 {
			goto st282
		}
		goto st0
	st282:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof282
		}
	st_case_282:
		if ( t.data)[( t.p)] == 108 {
			goto st283
		}
		goto st0
	st283:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof283
		}
	st_case_283:
		if ( t.data)[( t.p)] == 97 {
			goto st284
		}
		goto st0
	st284:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof284
		}
	st_case_284:
		if ( t.data)[( t.p)] == 121 {
			goto st239
		}
		goto st0
tr196:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st285
	st285:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof285
		}
	st_case_285:
//line tokenizer.go:5036
		if ( t.data)[( t.p)] == 97 {
			goto st286
		}
		goto st0
	st286:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof286
		}
	st_case_286:
		if ( t.data)[( t.p)] == 109 {
			goto st287
		}
		goto st0
	st287:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof287
		}
	st_case_287:
		if ( t.data)[( t.p)] == 97 {
			goto st288
		}
		goto st0
	st288:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof288
		}
	st_case_288:
		if ( t.data)[( t.p)] == 103 {
			goto st289
		}
		goto st0
	st289:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof289
		}
	st_case_289:
		if ( t.data)[( t.p)] == 101 {
			goto tr314
		}
		goto st0
tr314:
//line tokenizer.go.rl:121
 t.tok(DAMAGE) 
	goto st290
	st290:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof290
		}
	st_case_290:
//line tokenizer.go:5086
		if ( t.data)[( t.p)] == 32 {
			goto st291
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st291
		}
		goto st0
	st291:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof291
		}
	st_case_291:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st291
		case 40:
			goto tr316
		case 95:
			goto tr317
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st291
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr317
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr317
			}
		default:
			goto tr317
		}
		goto st0
tr316:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st292
	st292:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof292
		}
	st_case_292:
//line tokenizer.go:5139
		if ( t.data)[( t.p)] == 95 {
			goto st293
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st293
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st293
			}
		default:
			goto st293
		}
		goto st0
	st293:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof293
		}
	st_case_293:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st294
		case 95:
			goto st293
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st293
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st293
			}
		default:
			goto st293
		}
		goto st0
	st294:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof294
		}
	st_case_294:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr320
		case 124:
			goto tr321
		}
		goto st0
tr320:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st295
	st295:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof295
		}
	st_case_295:
//line tokenizer.go:5205
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr322
		case 95:
			goto tr323
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr323
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr323
			}
		default:
			goto tr323
		}
		goto st0
tr322:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st296
	st296:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof296
		}
	st_case_296:
//line tokenizer.go:5239
		if ( t.data)[( t.p)] == 95 {
			goto st297
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st297
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st297
			}
		default:
			goto st297
		}
		goto st0
	st297:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof297
		}
	st_case_297:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st298
		case 95:
			goto st297
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st297
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st297
			}
		default:
			goto st297
		}
		goto st0
	st298:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof298
		}
	st_case_298:
		if ( t.data)[( t.p)] == 41 {
			goto tr326
		}
		goto st0
tr326:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st299
	st299:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof299
		}
	st_case_299:
//line tokenizer.go:5302
		if ( t.data)[( t.p)] == 124 {
			goto tr327
		}
		goto st0
tr321:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st300
tr327:
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st300
	st300:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof300
		}
	st_case_300:
//line tokenizer.go:5324
		if ( t.data)[( t.p)] == 45 {
			goto st301
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr329
		}
		goto st0
	st301:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof301
		}
	st_case_301:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr329
		}
		goto st0
tr329:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st302
	st302:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof302
		}
	st_case_302:
//line tokenizer.go:5355
		if ( t.data)[( t.p)] == 32 {
			goto tr330
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st302
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr330
		}
		goto st0
tr330:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 303; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st303
	st303:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof303
		}
	st_case_303:
//line tokenizer.go:5384
		switch ( t.data)[( t.p)] {
		case 32:
			goto st303
		case 45:
			goto st304
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st303
		}
		goto st0
	st304:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof304
		}
	st_case_304:
		if ( t.data)[( t.p)] == 62 {
			goto tr334
		}
		goto st0
tr334:
//line tokenizer.go.rl:148
t.tok(ARROW)
	goto st305
	st305:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof305
		}
	st_case_305:
//line tokenizer.go:5413
		if ( t.data)[( t.p)] == 32 {
			goto st306
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st306
		}
		goto st0
	st306:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st306
		case 40:
			goto tr336
		case 95:
			goto tr337
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st306
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr337
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr337
			}
		default:
			goto tr337
		}
		goto st0
tr336:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st307
	st307:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof307
		}
	st_case_307:
//line tokenizer.go:5466
		if ( t.data)[( t.p)] == 95 {
			goto st308
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st308
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st308
			}
		default:
			goto st308
		}
		goto st0
	st308:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof308
		}
	st_case_308:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st309
		case 95:
			goto st308
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st308
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st308
			}
		default:
			goto st308
		}
		goto st0
	st309:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof309
		}
	st_case_309:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr340
		case 124:
			goto tr341
		}
		goto st0
tr340:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st310
	st310:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof310
		}
	st_case_310:
//line tokenizer.go:5532
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr342
		case 95:
			goto tr343
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr343
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr343
			}
		default:
			goto tr343
		}
		goto st0
tr342:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st311
	st311:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof311
		}
	st_case_311:
//line tokenizer.go:5566
		if ( t.data)[( t.p)] == 95 {
			goto st312
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st312
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st312
			}
		default:
			goto st312
		}
		goto st0
	st312:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st313
		case 95:
			goto st312
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st312
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st312
			}
		default:
			goto st312
		}
		goto st0
	st313:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof313
		}
	st_case_313:
		if ( t.data)[( t.p)] == 41 {
			goto tr346
		}
		goto st0
tr346:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st314
	st314:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof314
		}
	st_case_314:
//line tokenizer.go:5629
		if ( t.data)[( t.p)] == 124 {
			goto tr347
		}
		goto st0
tr341:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st315
tr347:
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st315
	st315:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof315
		}
	st_case_315:
//line tokenizer.go:5651
		if ( t.data)[( t.p)] == 45 {
			goto st316
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr349
		}
		goto st0
	st316:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof316
		}
	st_case_316:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr349
		}
		goto st0
tr349:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st317
	st317:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof317
		}
	st_case_317:
//line tokenizer.go:5682
		if ( t.data)[( t.p)] == 32 {
			goto tr350
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st317
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr350
		}
		goto st0
tr350:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 318; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st318
	st318:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof318
		}
	st_case_318:
//line tokenizer.go:5711
		switch ( t.data)[( t.p)] {
		case 32:
			goto st318
		case 45:
			goto st319
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st320
			}
		case ( t.data)[( t.p)] >= 9:
			goto st318
		}
		goto st0
	st319:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof319
		}
	st_case_319:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st320
		}
		goto st0
	st320:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st321
		case 46:
			goto st367
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st320
			}
		case ( t.data)[( t.p)] >= 9:
			goto st321
		}
		goto st0
	st321:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st321
		case 40:
			goto st322
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st321
		}
		goto st0
	st322:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof322
		}
	st_case_322:
		if ( t.data)[( t.p)] == 104 {
			goto st323
		}
		goto st0
	st323:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof323
		}
	st_case_323:
		if ( t.data)[( t.p)] == 58 {
			goto st324
		}
		goto st0
	st324:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof324
		}
	st_case_324:
		if ( t.data)[( t.p)] == 45 {
			goto st325
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st326
		}
		goto st0
	st325:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof325
		}
	st_case_325:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st326
		}
		goto st0
	st326:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st327
		case 46:
			goto st366
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st326
			}
		case ( t.data)[( t.p)] >= 9:
			goto st327
		}
		goto st0
	st327:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st327
		case 115:
			goto st328
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st327
		}
		goto st0
	st328:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof328
		}
	st_case_328:
		if ( t.data)[( t.p)] == 58 {
			goto st329
		}
		goto st0
	st329:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof329
		}
	st_case_329:
		if ( t.data)[( t.p)] == 45 {
			goto st330
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st331
		}
		goto st0
	st330:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof330
		}
	st_case_330:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st331
		}
		goto st0
	st331:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st332
		case 46:
			goto st365
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st331
		}
		goto st0
	st332:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof332
		}
	st_case_332:
		if ( t.data)[( t.p)] == 32 {
			goto st333
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st333
		}
		goto st0
	st333:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr371
		case 40:
			goto st335
		case 95:
			goto tr373
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr371
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr373
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr373
			}
		default:
			goto tr373
		}
		goto st0
tr371:
//line tokenizer.go.rl:140
t.tok(SOURCE)
	goto st334
	st334:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof334
		}
	st_case_334:
//line tokenizer.go:5942
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr371
		case 40:
			goto st335
		case 95:
			goto tr374
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr371
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr373
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr374
			}
		default:
			goto tr373
		}
		goto st0
	st335:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof335
		}
	st_case_335:
		if ( t.data)[( t.p)] == 95 {
			goto tr375
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr375
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr375
			}
		default:
			goto tr375
		}
		goto st0
tr375:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st336
	st336:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof336
		}
	st_case_336:
//line tokenizer.go:6004
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr376
		case 95:
			goto st336
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st336
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st336
			}
		default:
			goto st336
		}
		goto st0
tr376:
//line tokenizer.go.rl:138
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st337
	st337:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof337
		}
	st_case_337:
//line tokenizer.go:6033
		if ( t.data)[( t.p)] == 32 {
			goto st338
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st338
		}
		goto st0
tr406:
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st338
	st338:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof338
		}
	st_case_338:
//line tokenizer.go:6050
		switch ( t.data)[( t.p)] {
		case 32:
			goto st338
		case 95:
			goto tr379
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
				goto tr379
			}
		case ( t.data)[( t.p)] >= 9:
			goto st338
		}
		goto st0
tr379:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st339
	st339:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof339
		}
	st_case_339:
//line tokenizer.go:6080
		switch ( t.data)[( t.p)] {
		case 95:
			goto st339
		case 124:
			goto tr381
		}
		if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
			goto st339
		}
		goto st0
tr381:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st340
	st340:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof340
		}
	st_case_340:
//line tokenizer.go:6104
		if ( t.data)[( t.p)] == 95 {
			goto tr382
		}
		if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
			goto tr382
		}
		goto st0
tr382:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st684
	st684:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof684
		}
	st_case_684:
//line tokenizer.go:6126
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr761
		case 95:
			goto st684
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
				goto st684
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr761
		}
		goto st0
tr761:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st341
	st341:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof341
		}
	st_case_341:
//line tokenizer.go:6153
		switch ( t.data)[( t.p)] {
		case 32:
			goto st341
		case 60:
			goto st342
		case 82:
			goto tr385
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st341
		}
		goto st0
	st342:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof342
		}
	st_case_342:
		if ( t.data)[( t.p)] == 70 {
			goto st343
		}
		goto st0
	st343:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof343
		}
	st_case_343:
		if ( t.data)[( t.p)] == 114 {
			goto st344
		}
		goto st0
	st344:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof344
		}
	st_case_344:
		if ( t.data)[( t.p)] == 105 {
			goto st345
		}
		goto st0
	st345:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof345
		}
	st_case_345:
		if ( t.data)[( t.p)] == 101 {
			goto st346
		}
		goto st0
	st346:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof346
		}
	st_case_346:
		if ( t.data)[( t.p)] == 110 {
			goto st347
		}
		goto st0
	st347:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof347
		}
	st_case_347:
		if ( t.data)[( t.p)] == 100 {
			goto st348
		}
		goto st0
	st348:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof348
		}
	st_case_348:
		if ( t.data)[( t.p)] == 108 {
			goto st349
		}
		goto st0
	st349:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof349
		}
	st_case_349:
		if ( t.data)[( t.p)] == 121 {
			goto st350
		}
		goto st0
	st350:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof350
		}
	st_case_350:
		if ( t.data)[( t.p)] == 70 {
			goto st351
		}
		goto st0
	st351:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof351
		}
	st_case_351:
		if ( t.data)[( t.p)] == 105 {
			goto st352
		}
		goto st0
	st352:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof352
		}
	st_case_352:
		if ( t.data)[( t.p)] == 114 {
			goto st353
		}
		goto st0
	st353:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof353
		}
	st_case_353:
		if ( t.data)[( t.p)] == 101 {
			goto st354
		}
		goto st0
	st354:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof354
		}
	st_case_354:
		if ( t.data)[( t.p)] == 62 {
			goto tr398
		}
		goto st0
tr398:
//line tokenizer.go.rl:151
t.tok(FRIENDLY_FIRE)
	goto st685
	st685:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof685
		}
	st_case_685:
//line tokenizer.go:6292
		if ( t.data)[( t.p)] == 32 {
			goto st355
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st355
		}
		goto st0
	st355:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof355
		}
	st_case_355:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st355
		case 82:
			goto tr385
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st355
		}
		goto st0
tr385:
//line tokenizer.go.rl:152
t.tok(ROCKET)
	goto st356
	st356:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof356
		}
	st_case_356:
//line tokenizer.go:6324
		if ( t.data)[( t.p)] == 111 {
			goto st357
		}
		goto st0
	st357:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof357
		}
	st_case_357:
		if ( t.data)[( t.p)] == 99 {
			goto st358
		}
		goto st0
	st358:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof358
		}
	st_case_358:
		if ( t.data)[( t.p)] == 107 {
			goto st359
		}
		goto st0
	st359:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof359
		}
	st_case_359:
		if ( t.data)[( t.p)] == 101 {
			goto st360
		}
		goto st0
	st360:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof360
		}
	st_case_360:
		if ( t.data)[( t.p)] == 116 {
			goto st361
		}
		goto st0
	st361:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof361
		}
	st_case_361:
		if ( t.data)[( t.p)] == 32 {
			goto st362
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st362
		}
		goto st0
	st362:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof362
		}
	st_case_362:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st362
		case 45:
			goto st78
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr89
			}
		case ( t.data)[( t.p)] >= 9:
			goto st362
		}
		goto st0
tr373:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st363
	st363:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof363
		}
	st_case_363:
//line tokenizer.go:6411
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr406
		case 95:
			goto st363
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr406
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st363
				}
			case ( t.data)[( t.p)] >= 65:
				goto st363
			}
		default:
			goto st363
		}
		goto st0
tr374:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st364
	st364:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof364
		}
	st_case_364:
//line tokenizer.go:6450
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr406
		case 95:
			goto st364
		case 124:
			goto tr381
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr406
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st363
				}
			case ( t.data)[( t.p)] >= 65:
				goto st364
			}
		default:
			goto st363
		}
		goto st0
	st365:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof365
		}
	st_case_365:
		if ( t.data)[( t.p)] == 41 {
			goto st332
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st365
		}
		goto st0
	st366:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof366
		}
	st_case_366:
		if ( t.data)[( t.p)] == 32 {
			goto st327
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st366
			}
		case ( t.data)[( t.p)] >= 9:
			goto st327
		}
		goto st0
	st367:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof367
		}
	st_case_367:
		if ( t.data)[( t.p)] == 32 {
			goto st321
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st367
			}
		case ( t.data)[( t.p)] >= 9:
			goto st321
		}
		goto st0
tr343:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st368
	st368:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof368
		}
	st_case_368:
//line tokenizer.go:6537
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr346
		case 95:
			goto st368
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st368
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st368
			}
		default:
			goto st368
		}
		goto st0
tr337:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st369
	st369:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof369
		}
	st_case_369:
//line tokenizer.go:6571
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr340
		case 95:
			goto st369
		case 124:
			goto tr341
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st369
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st369
			}
		default:
			goto st369
		}
		goto st0
tr323:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st370
	st370:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof370
		}
	st_case_370:
//line tokenizer.go:6607
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr326
		case 95:
			goto st370
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st370
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st370
			}
		default:
			goto st370
		}
		goto st0
tr317:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st371
	st371:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof371
		}
	st_case_371:
//line tokenizer.go:6641
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr320
		case 95:
			goto st371
		case 124:
			goto tr321
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st371
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st371
			}
		default:
			goto st371
		}
		goto st0
tr197:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st372
	st372:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof372
		}
	st_case_372:
//line tokenizer.go:6677
		if ( t.data)[( t.p)] == 97 {
			goto st373
		}
		goto st0
	st373:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof373
		}
	st_case_373:
		if ( t.data)[( t.p)] == 109 {
			goto st374
		}
		goto st0
	st374:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof374
		}
	st_case_374:
		if ( t.data)[( t.p)] == 101 {
			goto st375
		}
		goto st0
	st375:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof375
		}
	st_case_375:
		if ( t.data)[( t.p)] == 112 {
			goto st376
		}
		goto st0
	st376:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof376
		}
	st_case_376:
		if ( t.data)[( t.p)] == 108 {
			goto st377
		}
		goto st0
	st377:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof377
		}
	st_case_377:
		if ( t.data)[( t.p)] == 97 {
			goto st378
		}
		goto st0
	st378:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof378
		}
	st_case_378:
		if ( t.data)[( t.p)] == 121 {
			goto st379
		}
		goto st0
	st379:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof379
		}
	st_case_379:
		if ( t.data)[( t.p)] == 32 {
			goto st380
		}
		goto st0
	st380:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof380
		}
	st_case_380:
		if ( t.data)[( t.p)] == 102 {
			goto st381
		}
		goto st0
	st381:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof381
		}
	st_case_381:
		if ( t.data)[( t.p)] == 105 {
			goto st382
		}
		goto st0
	st382:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof382
		}
	st_case_382:
		if ( t.data)[( t.p)] == 110 {
			goto st383
		}
		goto st0
	st383:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof383
		}
	st_case_383:
		if ( t.data)[( t.p)] == 105 {
			goto st384
		}
		goto st0
	st384:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof384
		}
	st_case_384:
		if ( t.data)[( t.p)] == 115 {
			goto st385
		}
		goto st0
	st385:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof385
		}
	st_case_385:
		if ( t.data)[( t.p)] == 104 {
			goto st386
		}
		goto st0
	st386:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof386
		}
	st_case_386:
		if ( t.data)[( t.p)] == 101 {
			goto st387
		}
		goto st0
	st387:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof387
		}
	st_case_387:
		if ( t.data)[( t.p)] == 100 {
			goto st388
		}
		goto st0
	st388:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof388
		}
	st_case_388:
		if ( t.data)[( t.p)] == 46 {
			goto tr429
		}
		goto st0
tr429:
//line tokenizer.go.rl:126
 t.tok(GAMEPLAY_FINISHED) 
	goto st389
	st389:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof389
		}
	st_case_389:
//line tokenizer.go:6835
		if ( t.data)[( t.p)] == 32 {
			goto st390
		}
		goto st0
	st390:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof390
		}
	st_case_390:
		if ( t.data)[( t.p)] == 87 {
			goto st391
		}
		goto st0
	st391:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof391
		}
	st_case_391:
		if ( t.data)[( t.p)] == 105 {
			goto st392
		}
		goto st0
	st392:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof392
		}
	st_case_392:
		if ( t.data)[( t.p)] == 110 {
			goto st393
		}
		goto st0
	st393:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof393
		}
	st_case_393:
		if ( t.data)[( t.p)] == 110 {
			goto st394
		}
		goto st0
	st394:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof394
		}
	st_case_394:
		if ( t.data)[( t.p)] == 101 {
			goto st395
		}
		goto st0
	st395:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof395
		}
	st_case_395:
		if ( t.data)[( t.p)] == 114 {
			goto st396
		}
		goto st0
	st396:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof396
		}
	st_case_396:
		if ( t.data)[( t.p)] == 32 {
			goto st397
		}
		goto st0
	st397:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof397
		}
	st_case_397:
		if ( t.data)[( t.p)] == 116 {
			goto st398
		}
		goto st0
	st398:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof398
		}
	st_case_398:
		if ( t.data)[( t.p)] == 101 {
			goto st399
		}
		goto st0
	st399:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof399
		}
	st_case_399:
		if ( t.data)[( t.p)] == 97 {
			goto st400
		}
		goto st0
	st400:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof400
		}
	st_case_400:
		if ( t.data)[( t.p)] == 109 {
			goto st401
		}
		goto st0
	st401:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof401
		}
	st_case_401:
		if ( t.data)[( t.p)] == 58 {
			goto st402
		}
		goto st0
	st402:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof402
		}
	st_case_402:
		if ( t.data)[( t.p)] == 32 {
			goto st403
		}
		goto st0
	st403:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof403
		}
	st_case_403:
		if ( t.data)[( t.p)] == 45 {
			goto st404
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr445
		}
		goto st0
	st404:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof404
		}
	st_case_404:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr445
		}
		goto st0
tr445:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st405
	st405:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof405
		}
	st_case_405:
//line tokenizer.go:6992
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr446
		case 46:
			goto tr447
		case 47:
			goto tr448
		case 95:
			goto tr448
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr449
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr448
			}
		default:
			goto tr448
		}
		goto st0
tr446:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 406; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st406
	st406:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof406
		}
	st_case_406:
//line tokenizer.go:7032
		if ( t.data)[( t.p)] == 95 {
			goto tr450
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr450
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr450
			}
		default:
			goto tr450
		}
		goto st0
tr450:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st407
	st407:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof407
		}
	st_case_407:
//line tokenizer.go:7063
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr451
		case 95:
			goto st407
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st407
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st407
			}
		default:
			goto st407
		}
		goto st0
tr451:
//line tokenizer.go.rl:138
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st408
	st408:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof408
		}
	st_case_408:
//line tokenizer.go:7092
		if ( t.data)[( t.p)] == 46 {
			goto st409
		}
		goto st0
tr502:
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st409
tr447:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 409; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
//line tokenizer.go.rl:140
t.tok(SOURCE)
	goto st409
tr504:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 409; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
//line tokenizer.go.rl:140
t.tok(SOURCE)
	goto st409
	st409:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof409
		}
	st_case_409:
//line tokenizer.go:7134
		if ( t.data)[( t.p)] == 32 {
			goto st410
		}
		goto st0
	st410:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof410
		}
	st_case_410:
		if ( t.data)[( t.p)] == 70 {
			goto st411
		}
		goto st0
	st411:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof411
		}
	st_case_411:
		if ( t.data)[( t.p)] == 105 {
			goto st412
		}
		goto st0
	st412:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof412
		}
	st_case_412:
		if ( t.data)[( t.p)] == 110 {
			goto st413
		}
		goto st0
	st413:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof413
		}
	st_case_413:
		if ( t.data)[( t.p)] == 105 {
			goto st414
		}
		goto st0
	st414:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof414
		}
	st_case_414:
		if ( t.data)[( t.p)] == 115 {
			goto st415
		}
		goto st0
	st415:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof415
		}
	st_case_415:
		if ( t.data)[( t.p)] == 104 {
			goto st416
		}
		goto st0
	st416:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof416
		}
	st_case_416:
		if ( t.data)[( t.p)] == 32 {
			goto st417
		}
		goto st0
	st417:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof417
		}
	st_case_417:
		if ( t.data)[( t.p)] == 114 {
			goto st418
		}
		goto st0
	st418:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof418
		}
	st_case_418:
		if ( t.data)[( t.p)] == 101 {
			goto st419
		}
		goto st0
	st419:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof419
		}
	st_case_419:
		if ( t.data)[( t.p)] == 97 {
			goto st420
		}
		goto st0
	st420:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof420
		}
	st_case_420:
		if ( t.data)[( t.p)] == 115 {
			goto st421
		}
		goto st0
	st421:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof421
		}
	st_case_421:
		if ( t.data)[( t.p)] == 111 {
			goto st422
		}
		goto st0
	st422:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof422
		}
	st_case_422:
		if ( t.data)[( t.p)] == 110 {
			goto st423
		}
		goto st0
	st423:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof423
		}
	st_case_423:
		if ( t.data)[( t.p)] == 58 {
			goto st424
		}
		goto st0
	st424:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof424
		}
	st_case_424:
		if ( t.data)[( t.p)] == 32 {
			goto st425
		}
		goto st0
	st425:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof425
		}
	st_case_425:
		if ( t.data)[( t.p)] == 39 {
			goto st426
		}
		goto st0
	st426:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof426
		}
	st_case_426:
		if ( t.data)[( t.p)] == 95 {
			goto st427
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st427
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st427
			}
		default:
			goto st427
		}
		goto st0
	st427:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof427
		}
	st_case_427:
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr472
		case 39:
			goto tr473
		case 95:
			goto st427
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st427
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st427
			}
		default:
			goto st427
		}
		goto st0
tr472:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st428
	st428:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof428
		}
	st_case_428:
//line tokenizer.go:7344
		if ( t.data)[( t.p)] == 95 {
			goto st429
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st429
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st429
			}
		default:
			goto st429
		}
		goto st0
	st429:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof429
		}
	st_case_429:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st428
		case 39:
			goto tr476
		case 95:
			goto st429
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st429
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st429
			}
		default:
			goto st429
		}
		goto st0
tr476:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st430
tr473:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st430
	st430:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof430
		}
	st_case_430:
//line tokenizer.go:7411
		if ( t.data)[( t.p)] == 46 {
			goto st431
		}
		goto st0
	st431:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof431
		}
	st_case_431:
		if ( t.data)[( t.p)] == 32 {
			goto st432
		}
		goto st0
	st432:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof432
		}
	st_case_432:
		if ( t.data)[( t.p)] == 65 {
			goto st433
		}
		goto st0
	st433:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof433
		}
	st_case_433:
		if ( t.data)[( t.p)] == 99 {
			goto st434
		}
		goto st0
	st434:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof434
		}
	st_case_434:
		if ( t.data)[( t.p)] == 116 {
			goto st435
		}
		goto st0
	st435:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof435
		}
	st_case_435:
		if ( t.data)[( t.p)] == 117 {
			goto st436
		}
		goto st0
	st436:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof436
		}
	st_case_436:
		if ( t.data)[( t.p)] == 97 {
			goto st437
		}
		goto st0
	st437:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof437
		}
	st_case_437:
		if ( t.data)[( t.p)] == 108 {
			goto st438
		}
		goto st0
	st438:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof438
		}
	st_case_438:
		if ( t.data)[( t.p)] == 32 {
			goto st439
		}
		goto st0
	st439:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof439
		}
	st_case_439:
		if ( t.data)[( t.p)] == 103 {
			goto st440
		}
		goto st0
	st440:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof440
		}
	st_case_440:
		if ( t.data)[( t.p)] == 97 {
			goto st441
		}
		goto st0
	st441:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof441
		}
	st_case_441:
		if ( t.data)[( t.p)] == 109 {
			goto st442
		}
		goto st0
	st442:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof442
		}
	st_case_442:
		if ( t.data)[( t.p)] == 101 {
			goto st443
		}
		goto st0
	st443:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof443
		}
	st_case_443:
		if ( t.data)[( t.p)] == 32 {
			goto st444
		}
		goto st0
	st444:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof444
		}
	st_case_444:
		if ( t.data)[( t.p)] == 116 {
			goto st445
		}
		goto st0
	st445:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof445
		}
	st_case_445:
		if ( t.data)[( t.p)] == 105 {
			goto st446
		}
		goto st0
	st446:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof446
		}
	st_case_446:
		if ( t.data)[( t.p)] == 109 {
			goto st447
		}
		goto st0
	st447:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof447
		}
	st_case_447:
		if ( t.data)[( t.p)] == 101 {
			goto st448
		}
		goto st0
	st448:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof448
		}
	st_case_448:
		if ( t.data)[( t.p)] == 32 {
			goto st449
		}
		goto st0
	st449:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof449
		}
	st_case_449:
		if ( t.data)[( t.p)] == 45 {
			goto st450
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st451
		}
		goto st0
	st450:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof450
		}
	st_case_450:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st451
		}
		goto st0
	st451:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof451
		}
	st_case_451:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st452
		case 46:
			goto st455
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st451
		}
		goto st0
	st452:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof452
		}
	st_case_452:
		if ( t.data)[( t.p)] == 115 {
			goto st453
		}
		goto st0
	st453:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof453
		}
	st_case_453:
		if ( t.data)[( t.p)] == 101 {
			goto st454
		}
		goto st0
	st454:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof454
		}
	st_case_454:
		if ( t.data)[( t.p)] == 99 {
			goto st680
		}
		goto st0
	st455:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof455
		}
	st_case_455:
		if ( t.data)[( t.p)] == 32 {
			goto st452
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st455
		}
		goto st0
tr448:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 456; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st456
	st456:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof456
		}
	st_case_456:
//line tokenizer.go:7676
		switch ( t.data)[( t.p)] {
		case 46:
			goto tr502
		case 95:
			goto st456
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st456
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st456
			}
		default:
			goto st456
		}
		goto st0
tr449:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 457; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st457
	st457:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof457
		}
	st_case_457:
//line tokenizer.go:7719
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr446
		case 46:
			goto tr504
		case 47:
			goto tr448
		case 95:
			goto tr448
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr449
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr448
			}
		default:
			goto tr448
		}
		goto st0
tr198:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st458
	st458:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof458
		}
	st_case_458:
//line tokenizer.go:7757
		if ( t.data)[( t.p)] == 101 {
			goto st459
		}
		goto st0
	st459:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof459
		}
	st_case_459:
		if ( t.data)[( t.p)] == 97 {
			goto st460
		}
		goto st0
	st460:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof460
		}
	st_case_460:
		if ( t.data)[( t.p)] == 108 {
			goto tr507
		}
		goto st0
tr507:
//line tokenizer.go.rl:122
 t.tok(HEAL) 
	goto st461
	st461:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof461
		}
	st_case_461:
//line tokenizer.go:7789
		if ( t.data)[( t.p)] == 32 {
			goto st462
		}
		goto st0
	st462:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof462
		}
	st_case_462:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st462
		case 40:
			goto tr509
		case 95:
			goto tr510
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr510
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr510
			}
		default:
			goto tr510
		}
		goto st0
tr509:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st463
	st463:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof463
		}
	st_case_463:
//line tokenizer.go:7834
		if ( t.data)[( t.p)] == 95 {
			goto st464
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st464
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st464
			}
		default:
			goto st464
		}
		goto st0
	st464:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof464
		}
	st_case_464:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st465
		case 95:
			goto st464
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st464
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st464
			}
		default:
			goto st464
		}
		goto st0
	st465:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof465
		}
	st_case_465:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr513
		case 124:
			goto tr514
		}
		goto st0
tr513:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st466
	st466:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof466
		}
	st_case_466:
//line tokenizer.go:7900
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr515
		case 95:
			goto tr516
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr516
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr516
			}
		default:
			goto tr516
		}
		goto st0
tr515:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st467
	st467:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof467
		}
	st_case_467:
//line tokenizer.go:7934
		if ( t.data)[( t.p)] == 95 {
			goto st468
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st468
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st468
			}
		default:
			goto st468
		}
		goto st0
	st468:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof468
		}
	st_case_468:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st469
		case 95:
			goto st468
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st468
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st468
			}
		default:
			goto st468
		}
		goto st0
	st469:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof469
		}
	st_case_469:
		if ( t.data)[( t.p)] == 41 {
			goto tr519
		}
		goto st0
tr519:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st470
	st470:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof470
		}
	st_case_470:
//line tokenizer.go:7997
		if ( t.data)[( t.p)] == 124 {
			goto tr520
		}
		goto st0
tr514:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st471
tr520:
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st471
	st471:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof471
		}
	st_case_471:
//line tokenizer.go:8019
		if ( t.data)[( t.p)] == 45 {
			goto st472
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr522
		}
		goto st0
	st472:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof472
		}
	st_case_472:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr522
		}
		goto st0
tr522:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st473
	st473:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof473
		}
	st_case_473:
//line tokenizer.go:8050
		if ( t.data)[( t.p)] == 32 {
			goto tr523
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st473
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr523
		}
		goto st0
tr523:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 474; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st474
	st474:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof474
		}
	st_case_474:
//line tokenizer.go:8079
		switch ( t.data)[( t.p)] {
		case 32:
			goto st474
		case 45:
			goto st475
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st474
		}
		goto st0
	st475:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof475
		}
	st_case_475:
		if ( t.data)[( t.p)] == 62 {
			goto tr527
		}
		goto st0
tr527:
//line tokenizer.go.rl:148
t.tok(ARROW)
	goto st476
	st476:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof476
		}
	st_case_476:
//line tokenizer.go:8108
		if ( t.data)[( t.p)] == 32 {
			goto st477
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st477
		}
		goto st0
	st477:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof477
		}
	st_case_477:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st477
		case 40:
			goto tr529
		case 95:
			goto tr530
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st477
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr530
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr530
			}
		default:
			goto tr530
		}
		goto st0
tr529:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st478
	st478:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof478
		}
	st_case_478:
//line tokenizer.go:8161
		if ( t.data)[( t.p)] == 95 {
			goto st479
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st479
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st479
			}
		default:
			goto st479
		}
		goto st0
	st479:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof479
		}
	st_case_479:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st480
		case 95:
			goto st479
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st479
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st479
			}
		default:
			goto st479
		}
		goto st0
	st480:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof480
		}
	st_case_480:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr533
		case 124:
			goto tr534
		}
		goto st0
tr533:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st481
	st481:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof481
		}
	st_case_481:
//line tokenizer.go:8227
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr535
		case 95:
			goto tr536
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr536
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr536
			}
		default:
			goto tr536
		}
		goto st0
tr535:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st482
	st482:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof482
		}
	st_case_482:
//line tokenizer.go:8261
		if ( t.data)[( t.p)] == 95 {
			goto st483
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st483
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st483
			}
		default:
			goto st483
		}
		goto st0
	st483:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof483
		}
	st_case_483:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st484
		case 95:
			goto st483
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st483
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st483
			}
		default:
			goto st483
		}
		goto st0
	st484:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof484
		}
	st_case_484:
		if ( t.data)[( t.p)] == 41 {
			goto tr539
		}
		goto st0
tr539:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st485
	st485:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof485
		}
	st_case_485:
//line tokenizer.go:8324
		if ( t.data)[( t.p)] == 124 {
			goto tr540
		}
		goto st0
tr534:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st486
tr540:
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st486
	st486:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof486
		}
	st_case_486:
//line tokenizer.go:8346
		if ( t.data)[( t.p)] == 45 {
			goto st487
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr542
		}
		goto st0
	st487:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof487
		}
	st_case_487:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr542
		}
		goto st0
tr542:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st488
	st488:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof488
		}
	st_case_488:
//line tokenizer.go:8377
		if ( t.data)[( t.p)] == 32 {
			goto tr543
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st488
		}
		goto st0
tr543:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 489; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st489
	st489:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof489
		}
	st_case_489:
//line tokenizer.go:8401
		if ( t.data)[( t.p)] == 45 {
			goto st490
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st686
		}
		goto st0
	st490:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof490
		}
	st_case_490:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st686
		}
		goto st0
	st686:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof686
		}
	st_case_686:
		switch ( t.data)[( t.p)] {
		case 40:
			goto st491
		case 46:
			goto st687
		case 47:
			goto tr765
		case 95:
			goto tr765
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr766
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr765
			}
		default:
			goto tr765
		}
		goto st0
	st491:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof491
		}
	st_case_491:
		if ( t.data)[( t.p)] == 95 {
			goto tr547
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr547
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr547
			}
		default:
			goto tr547
		}
		goto st0
tr547:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st492
	st492:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof492
		}
	st_case_492:
//line tokenizer.go:8481
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr548
		case 95:
			goto st492
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st492
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st492
			}
		default:
			goto st492
		}
		goto st0
	st687:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof687
		}
	st_case_687:
		switch ( t.data)[( t.p)] {
		case 40:
			goto st491
		case 47:
			goto tr765
		case 95:
			goto tr765
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr767
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr765
			}
		default:
			goto tr765
		}
		goto st0
tr765:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st688
	st688:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof688
		}
	st_case_688:
//line tokenizer.go:8541
		if ( t.data)[( t.p)] == 95 {
			goto st688
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st688
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st688
			}
		default:
			goto st688
		}
		goto st0
tr767:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st689
	st689:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof689
		}
	st_case_689:
//line tokenizer.go:8572
		switch ( t.data)[( t.p)] {
		case 40:
			goto st491
		case 47:
			goto tr765
		case 95:
			goto tr765
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr767
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr765
			}
		default:
			goto tr765
		}
		goto st0
tr766:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st690
	st690:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof690
		}
	st_case_690:
//line tokenizer.go:8608
		switch ( t.data)[( t.p)] {
		case 40:
			goto st491
		case 46:
			goto st687
		case 47:
			goto tr765
		case 95:
			goto tr765
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr766
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr765
			}
		default:
			goto tr765
		}
		goto st0
tr536:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st493
	st493:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof493
		}
	st_case_493:
//line tokenizer.go:8646
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr539
		case 95:
			goto st493
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st493
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st493
			}
		default:
			goto st493
		}
		goto st0
tr530:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st494
	st494:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof494
		}
	st_case_494:
//line tokenizer.go:8680
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr533
		case 95:
			goto st494
		case 124:
			goto tr534
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st494
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st494
			}
		default:
			goto st494
		}
		goto st0
tr516:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st495
	st495:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof495
		}
	st_case_495:
//line tokenizer.go:8716
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr519
		case 95:
			goto st495
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st495
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st495
			}
		default:
			goto st495
		}
		goto st0
tr510:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st496
	st496:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof496
		}
	st_case_496:
//line tokenizer.go:8750
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr513
		case 95:
			goto st496
		case 124:
			goto tr514
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st496
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st496
			}
		default:
			goto st496
		}
		goto st0
tr199:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st497
	st497:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof497
		}
	st_case_497:
//line tokenizer.go:8786
		if ( t.data)[( t.p)] == 105 {
			goto st498
		}
		goto st0
	st498:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof498
		}
	st_case_498:
		if ( t.data)[( t.p)] == 108 {
			goto st499
		}
		goto st0
	st499:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof499
		}
	st_case_499:
		if ( t.data)[( t.p)] == 108 {
			goto st500
		}
		goto st0
	st500:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof500
		}
	st_case_500:
		if ( t.data)[( t.p)] == 101 {
			goto st501
		}
		goto st0
	st501:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof501
		}
	st_case_501:
		if ( t.data)[( t.p)] == 100 {
			goto tr558
		}
		goto st0
tr558:
//line tokenizer.go.rl:123
 t.tok(KILL) 
	goto st502
	st502:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof502
		}
	st_case_502:
//line tokenizer.go:8836
		if ( t.data)[( t.p)] == 32 {
			goto st503
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st503
		}
		goto st0
	st503:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof503
		}
	st_case_503:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr560
		case 95:
			goto tr561
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr561
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr560:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st504
	st504:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof504
		}
	st_case_504:
//line tokenizer.go:8882
		if ( t.data)[( t.p)] == 95 {
			goto st505
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st505
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st505
			}
		default:
			goto st505
		}
		goto st0
	st505:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof505
		}
	st_case_505:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st506
		case 95:
			goto st505
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st505
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st505
			}
		default:
			goto st505
		}
		goto st0
	st506:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof506
		}
	st_case_506:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr564
		case 124:
			goto tr565
		}
		goto st0
tr564:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st507
	st507:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof507
		}
	st_case_507:
//line tokenizer.go:8948
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr566
		case 95:
			goto tr567
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr567
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr567
			}
		default:
			goto tr567
		}
		goto st0
tr566:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st508
	st508:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof508
		}
	st_case_508:
//line tokenizer.go:8982
		if ( t.data)[( t.p)] == 95 {
			goto st509
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st509
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st509
			}
		default:
			goto st509
		}
		goto st0
	st509:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof509
		}
	st_case_509:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st510
		case 95:
			goto st509
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st509
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st509
			}
		default:
			goto st509
		}
		goto st0
	st510:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof510
		}
	st_case_510:
		if ( t.data)[( t.p)] == 41 {
			goto tr570
		}
		goto st0
tr570:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st511
	st511:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof511
		}
	st_case_511:
//line tokenizer.go:9045
		if ( t.data)[( t.p)] == 124 {
			goto tr571
		}
		goto st0
tr565:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st512
tr571:
//line tokenizer.go.rl:114
 t.tok(int(t.data[t.p]))
	goto st512
	st512:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof512
		}
	st_case_512:
//line tokenizer.go:9067
		if ( t.data)[( t.p)] == 45 {
			goto st513
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr573
		}
		goto st0
	st513:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof513
		}
	st_case_513:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr573
		}
		goto st0
tr573:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st514
	st514:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof514
		}
	st_case_514:
//line tokenizer.go:9098
		if ( t.data)[( t.p)] == 59 {
			goto tr575
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st514
		}
		goto st0
tr575:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 515; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st515
	st515:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof515
		}
	st_case_515:
//line tokenizer.go:9122
		if ( t.data)[( t.p)] == 32 {
			goto st516
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st516
		}
		goto st0
	st516:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof516
		}
	st_case_516:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st516
		case 107:
			goto st517
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st516
		}
		goto st0
	st517:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof517
		}
	st_case_517:
		if ( t.data)[( t.p)] == 105 {
			goto st518
		}
		goto st0
	st518:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof518
		}
	st_case_518:
		if ( t.data)[( t.p)] == 108 {
			goto st519
		}
		goto st0
	st519:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof519
		}
	st_case_519:
		if ( t.data)[( t.p)] == 108 {
			goto st520
		}
		goto st0
	st520:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof520
		}
	st_case_520:
		if ( t.data)[( t.p)] == 101 {
			goto st521
		}
		goto st0
	st521:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof521
		}
	st_case_521:
		if ( t.data)[( t.p)] == 114 {
			goto st522
		}
		goto st0
	st522:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof522
		}
	st_case_522:
		if ( t.data)[( t.p)] == 32 {
			goto st523
		}
		goto st0
	st523:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof523
		}
	st_case_523:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr584
		case 95:
			goto tr585
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr585
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr585
			}
		default:
			goto tr585
		}
		goto st0
tr584:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st524
	st524:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof524
		}
	st_case_524:
//line tokenizer.go:9237
		if ( t.data)[( t.p)] == 95 {
			goto st525
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st525
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st525
			}
		default:
			goto st525
		}
		goto st0
	st525:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof525
		}
	st_case_525:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st526
		case 95:
			goto st525
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st525
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st525
			}
		default:
			goto st525
		}
		goto st0
	st526:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof526
		}
	st_case_526:
		if ( t.data)[( t.p)] == 124 {
			goto tr588
		}
		goto st0
tr588:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st527
	st527:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof527
		}
	st_case_527:
//line tokenizer.go:9298
		if ( t.data)[( t.p)] == 45 {
			goto st528
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr590
		}
		goto st0
	st528:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof528
		}
	st_case_528:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr590
		}
		goto st0
tr590:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st529
	st529:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof529
		}
	st_case_529:
//line tokenizer.go:9329
		if ( t.data)[( t.p)] == 32 {
			goto tr591
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st529
		}
		goto st0
tr591:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 530; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st530
	st530:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof530
		}
	st_case_530:
//line tokenizer.go:9353
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr593
		case 95:
			goto tr594
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr594
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr594
			}
		default:
			goto tr594
		}
		goto st0
tr593:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st531
	st531:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof531
		}
	st_case_531:
//line tokenizer.go:9387
		if ( t.data)[( t.p)] == 95 {
			goto st532
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st532
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st532
			}
		default:
			goto st532
		}
		goto st0
	st532:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof532
		}
	st_case_532:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st691
		case 95:
			goto st532
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st532
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st532
			}
		default:
			goto st532
		}
		goto st0
	st691:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof691
		}
	st_case_691:
		if ( t.data)[( t.p)] == 32 {
			goto tr769
		}
		goto st0
tr769:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st533
	st533:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof533
		}
	st_case_533:
//line tokenizer.go:9448
		if ( t.data)[( t.p)] == 60 {
			goto st534
		}
		goto st0
	st534:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof534
		}
	st_case_534:
		if ( t.data)[( t.p)] == 70 {
			goto st535
		}
		goto st0
	st535:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof535
		}
	st_case_535:
		if ( t.data)[( t.p)] == 114 {
			goto st536
		}
		goto st0
	st536:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof536
		}
	st_case_536:
		if ( t.data)[( t.p)] == 105 {
			goto st537
		}
		goto st0
	st537:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof537
		}
	st_case_537:
		if ( t.data)[( t.p)] == 101 {
			goto st538
		}
		goto st0
	st538:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof538
		}
	st_case_538:
		if ( t.data)[( t.p)] == 110 {
			goto st539
		}
		goto st0
	st539:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof539
		}
	st_case_539:
		if ( t.data)[( t.p)] == 100 {
			goto st540
		}
		goto st0
	st540:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof540
		}
	st_case_540:
		if ( t.data)[( t.p)] == 108 {
			goto st541
		}
		goto st0
	st541:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof541
		}
	st_case_541:
		if ( t.data)[( t.p)] == 121 {
			goto st542
		}
		goto st0
	st542:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof542
		}
	st_case_542:
		if ( t.data)[( t.p)] == 70 {
			goto st543
		}
		goto st0
	st543:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof543
		}
	st_case_543:
		if ( t.data)[( t.p)] == 105 {
			goto st544
		}
		goto st0
	st544:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof544
		}
	st_case_544:
		if ( t.data)[( t.p)] == 114 {
			goto st545
		}
		goto st0
	st545:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof545
		}
	st_case_545:
		if ( t.data)[( t.p)] == 101 {
			goto st546
		}
		goto st0
	st546:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof546
		}
	st_case_546:
		if ( t.data)[( t.p)] == 62 {
			goto tr610
		}
		goto st0
tr594:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st692
	st692:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof692
		}
	st_case_692:
//line tokenizer.go:9584
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr769
		case 95:
			goto st692
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st692
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st692
			}
		default:
			goto st692
		}
		goto st0
tr585:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st547
	st547:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof547
		}
	st_case_547:
//line tokenizer.go:9618
		switch ( t.data)[( t.p)] {
		case 95:
			goto st547
		case 124:
			goto tr588
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st547
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st547
			}
		default:
			goto st547
		}
		goto st0
tr567:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st548
	st548:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof548
		}
	st_case_548:
//line tokenizer.go:9652
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr570
		case 95:
			goto st548
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st548
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st548
			}
		default:
			goto st548
		}
		goto st0
tr561:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st549
	st549:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof549
		}
	st_case_549:
//line tokenizer.go:9686
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr564
		case 95:
			goto st549
		case 124:
			goto tr565
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st549
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st549
			}
		default:
			goto st549
		}
		goto st0
tr200:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st550
	st550:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof550
		}
	st_case_550:
//line tokenizer.go:9722
		if ( t.data)[( t.p)] == 97 {
			goto st551
		}
		goto st0
	st551:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof551
		}
	st_case_551:
		if ( t.data)[( t.p)] == 114 {
			goto st552
		}
		goto st0
	st552:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof552
		}
	st_case_552:
		if ( t.data)[( t.p)] == 116 {
			goto st553
		}
		goto st0
	st553:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof553
		}
	st_case_553:
		if ( t.data)[( t.p)] == 105 {
			goto st554
		}
		goto st0
	st554:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof554
		}
	st_case_554:
		if ( t.data)[( t.p)] == 99 {
			goto st555
		}
		goto st0
	st555:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof555
		}
	st_case_555:
		if ( t.data)[( t.p)] == 105 {
			goto st556
		}
		goto st0
	st556:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof556
		}
	st_case_556:
		if ( t.data)[( t.p)] == 112 {
			goto st557
		}
		goto st0
	st557:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof557
		}
	st_case_557:
		if ( t.data)[( t.p)] == 97 {
			goto st558
		}
		goto st0
	st558:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof558
		}
	st_case_558:
		if ( t.data)[( t.p)] == 110 {
			goto st559
		}
		goto st0
	st559:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof559
		}
	st_case_559:
		if ( t.data)[( t.p)] == 116 {
			goto tr623
		}
		goto st0
tr623:
//line tokenizer.go.rl:124
 t.tok(PARTICIPANT) 
	goto st560
	st560:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof560
		}
	st_case_560:
//line tokenizer.go:9817
		if ( t.data)[( t.p)] == 32 {
			goto st561
		}
		goto st0
	st561:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof561
		}
	st_case_561:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st561
		case 40:
			goto tr625
		case 95:
			goto tr626
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr626
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr626
			}
		default:
			goto tr626
		}
		goto st0
tr625:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st562
	st562:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof562
		}
	st_case_562:
//line tokenizer.go:9862
		if ( t.data)[( t.p)] == 95 {
			goto st563
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st563
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st563
			}
		default:
			goto st563
		}
		goto st0
	st563:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof563
		}
	st_case_563:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st564
		case 95:
			goto st563
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st563
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st563
			}
		default:
			goto st563
		}
		goto st0
	st564:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof564
		}
	st_case_564:
		if ( t.data)[( t.p)] == 9 {
			goto tr629
		}
		goto st0
tr629:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st565
	st565:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof565
		}
	st_case_565:
//line tokenizer.go:9923
		if ( t.data)[( t.p)] == 32 {
			goto st566
		}
		goto st0
	st566:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof566
		}
	st_case_566:
		switch ( t.data)[( t.p)] {
		case 9:
			goto st693
		case 32:
			goto st615
		case 40:
			goto tr633
		case 95:
			goto tr634
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr634
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr634
			}
		default:
			goto tr634
		}
		goto st0
tr690:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st693
	st693:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof693
		}
	st_case_693:
//line tokenizer.go:9967
		switch ( t.data)[( t.p)] {
		case 32:
			goto st694
		case 60:
			goto st567
		case 116:
			goto st577
		}
		goto st0
	st694:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof694
		}
	st_case_694:
		switch ( t.data)[( t.p)] {
		case 60:
			goto st567
		case 116:
			goto st577
		}
		goto st0
	st567:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof567
		}
	st_case_567:
		switch ( t.data)[( t.p)] {
		case 70:
			goto st535
		case 98:
			goto tr635
		case 100:
			goto tr636
		case 104:
			goto tr637
		}
		goto st0
tr635:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st568
	st568:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof568
		}
	st_case_568:
//line tokenizer.go:10019
		if ( t.data)[( t.p)] == 117 {
			goto st569
		}
		goto st0
	st569:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof569
		}
	st_case_569:
		if ( t.data)[( t.p)] == 102 {
			goto st570
		}
		goto st0
	st570:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof570
		}
	st_case_570:
		if ( t.data)[( t.p)] == 102 {
			goto st571
		}
		goto st0
	st571:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof571
		}
	st_case_571:
		if ( t.data)[( t.p)] == 62 {
			goto tr641
		}
		goto st0
tr641:
//line tokenizer.go.rl:153
t.tokval(StrTok(t.data[t.prev:t.p]))
	goto st695
	st695:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof695
		}
	st_case_695:
//line tokenizer.go:10060
		switch ( t.data)[( t.p)] {
		case 32:
			goto st696
		case 60:
			goto st534
		}
		goto st0
	st696:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof696
		}
	st_case_696:
		if ( t.data)[( t.p)] == 60 {
			goto st567
		}
		goto st0
tr636:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st572
	st572:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof572
		}
	st_case_572:
//line tokenizer.go:10091
		if ( t.data)[( t.p)] == 101 {
			goto st573
		}
		goto st0
	st573:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof573
		}
	st_case_573:
		if ( t.data)[( t.p)] == 98 {
			goto st568
		}
		goto st0
tr637:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st574
	st574:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof574
		}
	st_case_574:
//line tokenizer.go:10119
		if ( t.data)[( t.p)] == 101 {
			goto st575
		}
		goto st0
	st575:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof575
		}
	st_case_575:
		if ( t.data)[( t.p)] == 97 {
			goto st576
		}
		goto st0
	st576:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof576
		}
	st_case_576:
		if ( t.data)[( t.p)] == 108 {
			goto st571
		}
		goto st0
	st577:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof577
		}
	st_case_577:
		if ( t.data)[( t.p)] == 111 {
			goto st578
		}
		goto st0
	st578:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof578
		}
	st_case_578:
		if ( t.data)[( t.p)] == 116 {
			goto st579
		}
		goto st0
	st579:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof579
		}
	st_case_579:
		if ( t.data)[( t.p)] == 97 {
			goto st580
		}
		goto st0
	st580:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof580
		}
	st_case_580:
		if ( t.data)[( t.p)] == 108 {
			goto st581
		}
		goto st0
	st581:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof581
		}
	st_case_581:
		if ( t.data)[( t.p)] == 68 {
			goto st582
		}
		goto st0
	st582:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof582
		}
	st_case_582:
		if ( t.data)[( t.p)] == 97 {
			goto st583
		}
		goto st0
	st583:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof583
		}
	st_case_583:
		if ( t.data)[( t.p)] == 109 {
			goto st584
		}
		goto st0
	st584:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof584
		}
	st_case_584:
		if ( t.data)[( t.p)] == 97 {
			goto st585
		}
		goto st0
	st585:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof585
		}
	st_case_585:
		if ( t.data)[( t.p)] == 103 {
			goto st586
		}
		goto st0
	st586:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof586
		}
	st_case_586:
		if ( t.data)[( t.p)] == 101 {
			goto st587
		}
		goto st0
	st587:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof587
		}
	st_case_587:
		if ( t.data)[( t.p)] == 32 {
			goto st588
		}
		goto st0
	st588:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof588
		}
	st_case_588:
		if ( t.data)[( t.p)] == 45 {
			goto st589
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st590
		}
		goto st0
	st589:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof589
		}
	st_case_589:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st590
		}
		goto st0
	st590:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof590
		}
	st_case_590:
		switch ( t.data)[( t.p)] {
		case 46:
			goto st591
		case 59:
			goto st592
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st590
		}
		goto st0
	st591:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof591
		}
	st_case_591:
		if ( t.data)[( t.p)] == 59 {
			goto st592
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st591
		}
		goto st0
	st592:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof592
		}
	st_case_592:
		if ( t.data)[( t.p)] == 32 {
			goto st593
		}
		goto st0
	st593:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof593
		}
	st_case_593:
		if ( t.data)[( t.p)] == 109 {
			goto st594
		}
		goto st0
	st594:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof594
		}
	st_case_594:
		if ( t.data)[( t.p)] == 111 {
			goto st595
		}
		goto st0
	st595:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof595
		}
	st_case_595:
		if ( t.data)[( t.p)] == 115 {
			goto st596
		}
		goto st0
	st596:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof596
		}
	st_case_596:
		if ( t.data)[( t.p)] == 116 {
			goto st597
		}
		goto st0
	st597:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof597
		}
	st_case_597:
		if ( t.data)[( t.p)] == 68 {
			goto st598
		}
		goto st0
	st598:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof598
		}
	st_case_598:
		if ( t.data)[( t.p)] == 97 {
			goto st599
		}
		goto st0
	st599:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof599
		}
	st_case_599:
		if ( t.data)[( t.p)] == 109 {
			goto st600
		}
		goto st0
	st600:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof600
		}
	st_case_600:
		if ( t.data)[( t.p)] == 97 {
			goto st601
		}
		goto st0
	st601:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof601
		}
	st_case_601:
		if ( t.data)[( t.p)] == 103 {
			goto st602
		}
		goto st0
	st602:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof602
		}
	st_case_602:
		if ( t.data)[( t.p)] == 101 {
			goto st603
		}
		goto st0
	st603:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof603
		}
	st_case_603:
		if ( t.data)[( t.p)] == 87 {
			goto st604
		}
		goto st0
	st604:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof604
		}
	st_case_604:
		if ( t.data)[( t.p)] == 105 {
			goto st605
		}
		goto st0
	st605:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof605
		}
	st_case_605:
		if ( t.data)[( t.p)] == 116 {
			goto st606
		}
		goto st0
	st606:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof606
		}
	st_case_606:
		if ( t.data)[( t.p)] == 104 {
			goto st607
		}
		goto st0
	st607:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof607
		}
	st_case_607:
		if ( t.data)[( t.p)] == 32 {
			goto st608
		}
		goto st0
	st608:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof608
		}
	st_case_608:
		if ( t.data)[( t.p)] == 39 {
			goto st609
		}
		goto st0
	st609:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof609
		}
	st_case_609:
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr678
		case 40:
			goto st611
		case 95:
			goto tr680
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr680
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr680
			}
		default:
			goto tr680
		}
		goto st0
tr678:
//line tokenizer.go.rl:140
t.tok(SOURCE)
	goto st610
tr686:
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st610
	st610:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof610
		}
	st_case_610:
//line tokenizer.go:10481
		if ( t.data)[( t.p)] == 59 {
			goto st697
		}
		goto st0
	st697:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof697
		}
	st_case_697:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st696
		case 60:
			goto st567
		}
		goto st0
	st611:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof611
		}
	st_case_611:
		if ( t.data)[( t.p)] == 95 {
			goto tr682
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr682
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr682
			}
		default:
			goto tr682
		}
		goto st0
tr682:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st612
	st612:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof612
		}
	st_case_612:
//line tokenizer.go:10533
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr683
		case 95:
			goto st612
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st612
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st612
			}
		default:
			goto st612
		}
		goto st0
tr683:
//line tokenizer.go.rl:138
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
	goto st613
	st613:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof613
		}
	st_case_613:
//line tokenizer.go:10562
		if ( t.data)[( t.p)] == 39 {
			goto st610
		}
		goto st0
tr680:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st614
	st614:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof614
		}
	st_case_614:
//line tokenizer.go:10581
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr686
		case 95:
			goto st614
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st614
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st614
			}
		default:
			goto st614
		}
		goto st0
tr691:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st615
	st615:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof615
		}
	st_case_615:
//line tokenizer.go:10612
		switch ( t.data)[( t.p)] {
		case 9:
			goto st693
		case 32:
			goto st615
		}
		goto st0
tr633:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st616
	st616:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof616
		}
	st_case_616:
//line tokenizer.go:10634
		if ( t.data)[( t.p)] == 95 {
			goto st617
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st617
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st617
			}
		default:
			goto st617
		}
		goto st0
	st617:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof617
		}
	st_case_617:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st618
		case 95:
			goto st617
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st617
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st617
			}
		default:
			goto st617
		}
		goto st0
	st618:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof618
		}
	st_case_618:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr690
		case 32:
			goto tr691
		}
		goto st0
tr634:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st619
	st619:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof619
		}
	st_case_619:
//line tokenizer.go:10701
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr690
		case 32:
			goto tr691
		case 95:
			goto st619
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st619
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st619
			}
		default:
			goto st619
		}
		goto st0
tr626:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st620
	st620:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof620
		}
	st_case_620:
//line tokenizer.go:10737
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr629
		case 95:
			goto st620
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st620
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st620
			}
		default:
			goto st620
		}
		goto st0
tr201:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st621
	st621:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof621
		}
	st_case_621:
//line tokenizer.go:10771
		if ( t.data)[( t.p)] == 101 {
			goto st622
		}
		goto st0
	st622:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof622
		}
	st_case_622:
		if ( t.data)[( t.p)] == 119 {
			goto st623
		}
		goto st0
	st623:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof623
		}
	st_case_623:
		if ( t.data)[( t.p)] == 97 {
			goto st624
		}
		goto st0
	st624:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof624
		}
	st_case_624:
		if ( t.data)[( t.p)] == 114 {
			goto st625
		}
		goto st0
	st625:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof625
		}
	st_case_625:
		if ( t.data)[( t.p)] == 100 {
			goto tr698
		}
		goto st0
tr698:
//line tokenizer.go.rl:127
 t.tok(REWARD) 
	goto st626
	st626:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof626
		}
	st_case_626:
//line tokenizer.go:10821
		if ( t.data)[( t.p)] == 32 {
			goto st627
		}
		goto st0
	st627:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof627
		}
	st_case_627:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st627
		case 40:
			goto tr700
		case 95:
			goto tr701
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr701
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr701
			}
		default:
			goto tr701
		}
		goto st0
tr700:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st628
	st628:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof628
		}
	st_case_628:
//line tokenizer.go:10866
		if ( t.data)[( t.p)] == 95 {
			goto st629
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st629
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st629
			}
		default:
			goto st629
		}
		goto st0
	st629:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof629
		}
	st_case_629:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st630
		case 95:
			goto st629
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st629
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st629
			}
		default:
			goto st629
		}
		goto st0
	st630:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof630
		}
	st_case_630:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr704
		case 32:
			goto tr705
		}
		goto st0
tr704:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st631
	st631:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof631
		}
	st_case_631:
//line tokenizer.go:10930
		switch ( t.data)[( t.p)] {
		case 32:
			goto st631
		case 45:
			goto st632
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr708
		}
		goto st0
	st632:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof632
		}
	st_case_632:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr708
		}
		goto st0
tr708:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st633
	st633:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof633
		}
	st_case_633:
//line tokenizer.go:10964
		if ( t.data)[( t.p)] == 32 {
			goto tr709
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st633
		}
		goto st0
tr709:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 634; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
	goto st634
	st634:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof634
		}
	st_case_634:
//line tokenizer.go:10988
		switch ( t.data)[( t.p)] {
		case 99:
			goto st635
		case 101:
			goto st646
		case 107:
			goto st667
		}
		goto st0
	st635:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof635
		}
	st_case_635:
		if ( t.data)[( t.p)] == 114 {
			goto st636
		}
		goto st0
	st636:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof636
		}
	st_case_636:
		if ( t.data)[( t.p)] == 101 {
			goto st637
		}
		goto st0
	st637:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof637
		}
	st_case_637:
		if ( t.data)[( t.p)] == 100 {
			goto st638
		}
		goto st0
	st638:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof638
		}
	st_case_638:
		if ( t.data)[( t.p)] == 105 {
			goto st639
		}
		goto st0
	st639:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof639
		}
	st_case_639:
		if ( t.data)[( t.p)] == 116 {
			goto st640
		}
		goto st0
	st640:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof640
		}
	st_case_640:
		if ( t.data)[( t.p)] == 115 {
			goto st641
		}
		goto st0
	st641:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof641
		}
	st_case_641:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st641
		case 102:
			goto st642
		}
		goto st0
	st642:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof642
		}
	st_case_642:
		if ( t.data)[( t.p)] == 111 {
			goto st643
		}
		goto st0
	st643:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof643
		}
	st_case_643:
		if ( t.data)[( t.p)] == 114 {
			goto st644
		}
		goto st0
	st644:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof644
		}
	st_case_644:
		if ( t.data)[( t.p)] == 32 {
			goto st645
		}
		goto st0
	st645:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof645
		}
	st_case_645:
		if ( t.data)[( t.p)] == 10 {
			goto st0
		}
		goto tr724
	st646:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof646
		}
	st_case_646:
		switch ( t.data)[( t.p)] {
		case 102:
			goto st647
		case 120:
			goto st659
		}
		goto st0
	st647:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof647
		}
	st_case_647:
		if ( t.data)[( t.p)] == 102 {
			goto st648
		}
		goto st0
	st648:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof648
		}
	st_case_648:
		if ( t.data)[( t.p)] == 101 {
			goto st649
		}
		goto st0
	st649:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof649
		}
	st_case_649:
		if ( t.data)[( t.p)] == 99 {
			goto st650
		}
		goto st0
	st650:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof650
		}
	st_case_650:
		if ( t.data)[( t.p)] == 116 {
			goto st651
		}
		goto st0
	st651:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof651
		}
	st_case_651:
		if ( t.data)[( t.p)] == 105 {
			goto st652
		}
		goto st0
	st652:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof652
		}
	st_case_652:
		if ( t.data)[( t.p)] == 118 {
			goto st653
		}
		goto st0
	st653:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof653
		}
	st_case_653:
		if ( t.data)[( t.p)] == 101 {
			goto st654
		}
		goto st0
	st654:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof654
		}
	st_case_654:
		if ( t.data)[( t.p)] == 32 {
			goto st655
		}
		goto st0
	st655:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof655
		}
	st_case_655:
		if ( t.data)[( t.p)] == 112 {
			goto st656
		}
		goto st0
	st656:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof656
		}
	st_case_656:
		if ( t.data)[( t.p)] == 111 {
			goto st657
		}
		goto st0
	st657:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof657
		}
	st_case_657:
		if ( t.data)[( t.p)] == 105 {
			goto st658
		}
		goto st0
	st658:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof658
		}
	st_case_658:
		if ( t.data)[( t.p)] == 110 {
			goto st639
		}
		goto st0
	st659:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof659
		}
	st_case_659:
		if ( t.data)[( t.p)] == 112 {
			goto st660
		}
		goto st0
	st660:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof660
		}
	st_case_660:
		if ( t.data)[( t.p)] == 101 {
			goto st661
		}
		goto st0
	st661:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof661
		}
	st_case_661:
		if ( t.data)[( t.p)] == 114 {
			goto st662
		}
		goto st0
	st662:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof662
		}
	st_case_662:
		if ( t.data)[( t.p)] == 105 {
			goto st663
		}
		goto st0
	st663:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof663
		}
	st_case_663:
		if ( t.data)[( t.p)] == 101 {
			goto st664
		}
		goto st0
	st664:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof664
		}
	st_case_664:
		if ( t.data)[( t.p)] == 110 {
			goto st665
		}
		goto st0
	st665:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof665
		}
	st_case_665:
		if ( t.data)[( t.p)] == 99 {
			goto st666
		}
		goto st0
	st666:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof666
		}
	st_case_666:
		if ( t.data)[( t.p)] == 101 {
			goto st641
		}
		goto st0
	st667:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof667
		}
	st_case_667:
		if ( t.data)[( t.p)] == 97 {
			goto st668
		}
		goto st0
	st668:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof668
		}
	st_case_668:
		if ( t.data)[( t.p)] == 114 {
			goto st669
		}
		goto st0
	st669:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof669
		}
	st_case_669:
		if ( t.data)[( t.p)] == 109 {
			goto st670
		}
		goto st0
	st670:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof670
		}
	st_case_670:
		if ( t.data)[( t.p)] == 97 {
			goto st641
		}
		goto st0
tr705:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st671
	st671:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof671
		}
	st_case_671:
//line tokenizer.go:11339
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr748
		case 95:
			goto tr749
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr749
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr749
			}
		default:
			goto tr749
		}
		goto st0
tr748:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st672
	st672:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof672
		}
	st_case_672:
//line tokenizer.go:11373
		if ( t.data)[( t.p)] == 95 {
			goto st673
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st673
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st673
			}
		default:
			goto st673
		}
		goto st0
	st673:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof673
		}
	st_case_673:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st674
		case 95:
			goto st673
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st673
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st673
			}
		default:
			goto st673
		}
		goto st0
	st674:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof674
		}
	st_case_674:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr704
		case 32:
			goto tr752
		}
		goto st0
tr752:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
	goto st675
	st675:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof675
		}
	st_case_675:
//line tokenizer.go:11437
		switch ( t.data)[( t.p)] {
		case 9:
			goto st631
		case 32:
			goto st675
		}
		goto st0
tr749:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st676
	st676:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof676
		}
	st_case_676:
//line tokenizer.go:11459
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr704
		case 32:
			goto tr752
		case 95:
			goto st676
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st676
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st676
			}
		default:
			goto st676
		}
		goto st0
tr701:
//line tokenizer.go.rl:103

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st677
	st677:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof677
		}
	st_case_677:
//line tokenizer.go:11495
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr704
		case 32:
			goto tr705
		case 95:
			goto st677
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st677
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st677
			}
		default:
			goto st677
		}
		goto st0
	st_out:
	_test_eof2:  t.cs = 2; goto _test_eof
	_test_eof3:  t.cs = 3; goto _test_eof
	_test_eof4:  t.cs = 4; goto _test_eof
	_test_eof5:  t.cs = 5; goto _test_eof
	_test_eof6:  t.cs = 6; goto _test_eof
	_test_eof7:  t.cs = 7; goto _test_eof
	_test_eof8:  t.cs = 8; goto _test_eof
	_test_eof9:  t.cs = 9; goto _test_eof
	_test_eof10:  t.cs = 10; goto _test_eof
	_test_eof11:  t.cs = 11; goto _test_eof
	_test_eof12:  t.cs = 12; goto _test_eof
	_test_eof13:  t.cs = 13; goto _test_eof
	_test_eof14:  t.cs = 14; goto _test_eof
	_test_eof15:  t.cs = 15; goto _test_eof
	_test_eof16:  t.cs = 16; goto _test_eof
	_test_eof17:  t.cs = 17; goto _test_eof
	_test_eof18:  t.cs = 18; goto _test_eof
	_test_eof19:  t.cs = 19; goto _test_eof
	_test_eof20:  t.cs = 20; goto _test_eof
	_test_eof21:  t.cs = 21; goto _test_eof
	_test_eof22:  t.cs = 22; goto _test_eof
	_test_eof23:  t.cs = 23; goto _test_eof
	_test_eof24:  t.cs = 24; goto _test_eof
	_test_eof25:  t.cs = 25; goto _test_eof
	_test_eof26:  t.cs = 26; goto _test_eof
	_test_eof27:  t.cs = 27; goto _test_eof
	_test_eof28:  t.cs = 28; goto _test_eof
	_test_eof29:  t.cs = 29; goto _test_eof
	_test_eof30:  t.cs = 30; goto _test_eof
	_test_eof31:  t.cs = 31; goto _test_eof
	_test_eof32:  t.cs = 32; goto _test_eof
	_test_eof33:  t.cs = 33; goto _test_eof
	_test_eof34:  t.cs = 34; goto _test_eof
	_test_eof35:  t.cs = 35; goto _test_eof
	_test_eof36:  t.cs = 36; goto _test_eof
	_test_eof37:  t.cs = 37; goto _test_eof
	_test_eof38:  t.cs = 38; goto _test_eof
	_test_eof39:  t.cs = 39; goto _test_eof
	_test_eof40:  t.cs = 40; goto _test_eof
	_test_eof41:  t.cs = 41; goto _test_eof
	_test_eof42:  t.cs = 42; goto _test_eof
	_test_eof43:  t.cs = 43; goto _test_eof
	_test_eof44:  t.cs = 44; goto _test_eof
	_test_eof45:  t.cs = 45; goto _test_eof
	_test_eof46:  t.cs = 46; goto _test_eof
	_test_eof47:  t.cs = 47; goto _test_eof
	_test_eof48:  t.cs = 48; goto _test_eof
	_test_eof49:  t.cs = 49; goto _test_eof
	_test_eof50:  t.cs = 50; goto _test_eof
	_test_eof51:  t.cs = 51; goto _test_eof
	_test_eof52:  t.cs = 52; goto _test_eof
	_test_eof53:  t.cs = 53; goto _test_eof
	_test_eof54:  t.cs = 54; goto _test_eof
	_test_eof55:  t.cs = 55; goto _test_eof
	_test_eof56:  t.cs = 56; goto _test_eof
	_test_eof57:  t.cs = 57; goto _test_eof
	_test_eof58:  t.cs = 58; goto _test_eof
	_test_eof59:  t.cs = 59; goto _test_eof
	_test_eof60:  t.cs = 60; goto _test_eof
	_test_eof61:  t.cs = 61; goto _test_eof
	_test_eof62:  t.cs = 62; goto _test_eof
	_test_eof63:  t.cs = 63; goto _test_eof
	_test_eof64:  t.cs = 64; goto _test_eof
	_test_eof65:  t.cs = 65; goto _test_eof
	_test_eof66:  t.cs = 66; goto _test_eof
	_test_eof67:  t.cs = 67; goto _test_eof
	_test_eof68:  t.cs = 68; goto _test_eof
	_test_eof69:  t.cs = 69; goto _test_eof
	_test_eof70:  t.cs = 70; goto _test_eof
	_test_eof678:  t.cs = 678; goto _test_eof
	_test_eof71:  t.cs = 71; goto _test_eof
	_test_eof72:  t.cs = 72; goto _test_eof
	_test_eof73:  t.cs = 73; goto _test_eof
	_test_eof74:  t.cs = 74; goto _test_eof
	_test_eof75:  t.cs = 75; goto _test_eof
	_test_eof76:  t.cs = 76; goto _test_eof
	_test_eof77:  t.cs = 77; goto _test_eof
	_test_eof78:  t.cs = 78; goto _test_eof
	_test_eof679:  t.cs = 679; goto _test_eof
	_test_eof79:  t.cs = 79; goto _test_eof
	_test_eof80:  t.cs = 80; goto _test_eof
	_test_eof81:  t.cs = 81; goto _test_eof
	_test_eof82:  t.cs = 82; goto _test_eof
	_test_eof83:  t.cs = 83; goto _test_eof
	_test_eof84:  t.cs = 84; goto _test_eof
	_test_eof85:  t.cs = 85; goto _test_eof
	_test_eof86:  t.cs = 86; goto _test_eof
	_test_eof87:  t.cs = 87; goto _test_eof
	_test_eof88:  t.cs = 88; goto _test_eof
	_test_eof89:  t.cs = 89; goto _test_eof
	_test_eof90:  t.cs = 90; goto _test_eof
	_test_eof91:  t.cs = 91; goto _test_eof
	_test_eof92:  t.cs = 92; goto _test_eof
	_test_eof93:  t.cs = 93; goto _test_eof
	_test_eof94:  t.cs = 94; goto _test_eof
	_test_eof95:  t.cs = 95; goto _test_eof
	_test_eof96:  t.cs = 96; goto _test_eof
	_test_eof97:  t.cs = 97; goto _test_eof
	_test_eof98:  t.cs = 98; goto _test_eof
	_test_eof99:  t.cs = 99; goto _test_eof
	_test_eof100:  t.cs = 100; goto _test_eof
	_test_eof101:  t.cs = 101; goto _test_eof
	_test_eof102:  t.cs = 102; goto _test_eof
	_test_eof103:  t.cs = 103; goto _test_eof
	_test_eof104:  t.cs = 104; goto _test_eof
	_test_eof105:  t.cs = 105; goto _test_eof
	_test_eof106:  t.cs = 106; goto _test_eof
	_test_eof107:  t.cs = 107; goto _test_eof
	_test_eof108:  t.cs = 108; goto _test_eof
	_test_eof109:  t.cs = 109; goto _test_eof
	_test_eof110:  t.cs = 110; goto _test_eof
	_test_eof111:  t.cs = 111; goto _test_eof
	_test_eof112:  t.cs = 112; goto _test_eof
	_test_eof113:  t.cs = 113; goto _test_eof
	_test_eof114:  t.cs = 114; goto _test_eof
	_test_eof115:  t.cs = 115; goto _test_eof
	_test_eof116:  t.cs = 116; goto _test_eof
	_test_eof117:  t.cs = 117; goto _test_eof
	_test_eof118:  t.cs = 118; goto _test_eof
	_test_eof119:  t.cs = 119; goto _test_eof
	_test_eof120:  t.cs = 120; goto _test_eof
	_test_eof121:  t.cs = 121; goto _test_eof
	_test_eof122:  t.cs = 122; goto _test_eof
	_test_eof123:  t.cs = 123; goto _test_eof
	_test_eof124:  t.cs = 124; goto _test_eof
	_test_eof125:  t.cs = 125; goto _test_eof
	_test_eof126:  t.cs = 126; goto _test_eof
	_test_eof127:  t.cs = 127; goto _test_eof
	_test_eof128:  t.cs = 128; goto _test_eof
	_test_eof129:  t.cs = 129; goto _test_eof
	_test_eof130:  t.cs = 130; goto _test_eof
	_test_eof131:  t.cs = 131; goto _test_eof
	_test_eof132:  t.cs = 132; goto _test_eof
	_test_eof133:  t.cs = 133; goto _test_eof
	_test_eof134:  t.cs = 134; goto _test_eof
	_test_eof680:  t.cs = 680; goto _test_eof
	_test_eof135:  t.cs = 135; goto _test_eof
	_test_eof136:  t.cs = 136; goto _test_eof
	_test_eof137:  t.cs = 137; goto _test_eof
	_test_eof138:  t.cs = 138; goto _test_eof
	_test_eof139:  t.cs = 139; goto _test_eof
	_test_eof140:  t.cs = 140; goto _test_eof
	_test_eof141:  t.cs = 141; goto _test_eof
	_test_eof142:  t.cs = 142; goto _test_eof
	_test_eof143:  t.cs = 143; goto _test_eof
	_test_eof144:  t.cs = 144; goto _test_eof
	_test_eof145:  t.cs = 145; goto _test_eof
	_test_eof146:  t.cs = 146; goto _test_eof
	_test_eof147:  t.cs = 147; goto _test_eof
	_test_eof148:  t.cs = 148; goto _test_eof
	_test_eof149:  t.cs = 149; goto _test_eof
	_test_eof150:  t.cs = 150; goto _test_eof
	_test_eof151:  t.cs = 151; goto _test_eof
	_test_eof152:  t.cs = 152; goto _test_eof
	_test_eof153:  t.cs = 153; goto _test_eof
	_test_eof154:  t.cs = 154; goto _test_eof
	_test_eof681:  t.cs = 681; goto _test_eof
	_test_eof682:  t.cs = 682; goto _test_eof
	_test_eof155:  t.cs = 155; goto _test_eof
	_test_eof156:  t.cs = 156; goto _test_eof
	_test_eof157:  t.cs = 157; goto _test_eof
	_test_eof158:  t.cs = 158; goto _test_eof
	_test_eof159:  t.cs = 159; goto _test_eof
	_test_eof160:  t.cs = 160; goto _test_eof
	_test_eof161:  t.cs = 161; goto _test_eof
	_test_eof162:  t.cs = 162; goto _test_eof
	_test_eof163:  t.cs = 163; goto _test_eof
	_test_eof164:  t.cs = 164; goto _test_eof
	_test_eof165:  t.cs = 165; goto _test_eof
	_test_eof166:  t.cs = 166; goto _test_eof
	_test_eof167:  t.cs = 167; goto _test_eof
	_test_eof168:  t.cs = 168; goto _test_eof
	_test_eof169:  t.cs = 169; goto _test_eof
	_test_eof170:  t.cs = 170; goto _test_eof
	_test_eof171:  t.cs = 171; goto _test_eof
	_test_eof172:  t.cs = 172; goto _test_eof
	_test_eof173:  t.cs = 173; goto _test_eof
	_test_eof174:  t.cs = 174; goto _test_eof
	_test_eof175:  t.cs = 175; goto _test_eof
	_test_eof176:  t.cs = 176; goto _test_eof
	_test_eof177:  t.cs = 177; goto _test_eof
	_test_eof178:  t.cs = 178; goto _test_eof
	_test_eof179:  t.cs = 179; goto _test_eof
	_test_eof180:  t.cs = 180; goto _test_eof
	_test_eof181:  t.cs = 181; goto _test_eof
	_test_eof182:  t.cs = 182; goto _test_eof
	_test_eof183:  t.cs = 183; goto _test_eof
	_test_eof184:  t.cs = 184; goto _test_eof
	_test_eof185:  t.cs = 185; goto _test_eof
	_test_eof186:  t.cs = 186; goto _test_eof
	_test_eof187:  t.cs = 187; goto _test_eof
	_test_eof188:  t.cs = 188; goto _test_eof
	_test_eof189:  t.cs = 189; goto _test_eof
	_test_eof190:  t.cs = 190; goto _test_eof
	_test_eof191:  t.cs = 191; goto _test_eof
	_test_eof192:  t.cs = 192; goto _test_eof
	_test_eof193:  t.cs = 193; goto _test_eof
	_test_eof194:  t.cs = 194; goto _test_eof
	_test_eof195:  t.cs = 195; goto _test_eof
	_test_eof196:  t.cs = 196; goto _test_eof
	_test_eof197:  t.cs = 197; goto _test_eof
	_test_eof198:  t.cs = 198; goto _test_eof
	_test_eof199:  t.cs = 199; goto _test_eof
	_test_eof200:  t.cs = 200; goto _test_eof
	_test_eof201:  t.cs = 201; goto _test_eof
	_test_eof202:  t.cs = 202; goto _test_eof
	_test_eof203:  t.cs = 203; goto _test_eof
	_test_eof204:  t.cs = 204; goto _test_eof
	_test_eof205:  t.cs = 205; goto _test_eof
	_test_eof206:  t.cs = 206; goto _test_eof
	_test_eof207:  t.cs = 207; goto _test_eof
	_test_eof208:  t.cs = 208; goto _test_eof
	_test_eof209:  t.cs = 209; goto _test_eof
	_test_eof210:  t.cs = 210; goto _test_eof
	_test_eof211:  t.cs = 211; goto _test_eof
	_test_eof212:  t.cs = 212; goto _test_eof
	_test_eof213:  t.cs = 213; goto _test_eof
	_test_eof214:  t.cs = 214; goto _test_eof
	_test_eof215:  t.cs = 215; goto _test_eof
	_test_eof216:  t.cs = 216; goto _test_eof
	_test_eof217:  t.cs = 217; goto _test_eof
	_test_eof218:  t.cs = 218; goto _test_eof
	_test_eof219:  t.cs = 219; goto _test_eof
	_test_eof220:  t.cs = 220; goto _test_eof
	_test_eof221:  t.cs = 221; goto _test_eof
	_test_eof222:  t.cs = 222; goto _test_eof
	_test_eof223:  t.cs = 223; goto _test_eof
	_test_eof224:  t.cs = 224; goto _test_eof
	_test_eof225:  t.cs = 225; goto _test_eof
	_test_eof226:  t.cs = 226; goto _test_eof
	_test_eof227:  t.cs = 227; goto _test_eof
	_test_eof228:  t.cs = 228; goto _test_eof
	_test_eof229:  t.cs = 229; goto _test_eof
	_test_eof230:  t.cs = 230; goto _test_eof
	_test_eof231:  t.cs = 231; goto _test_eof
	_test_eof232:  t.cs = 232; goto _test_eof
	_test_eof233:  t.cs = 233; goto _test_eof
	_test_eof234:  t.cs = 234; goto _test_eof
	_test_eof235:  t.cs = 235; goto _test_eof
	_test_eof236:  t.cs = 236; goto _test_eof
	_test_eof237:  t.cs = 237; goto _test_eof
	_test_eof238:  t.cs = 238; goto _test_eof
	_test_eof239:  t.cs = 239; goto _test_eof
	_test_eof240:  t.cs = 240; goto _test_eof
	_test_eof241:  t.cs = 241; goto _test_eof
	_test_eof242:  t.cs = 242; goto _test_eof
	_test_eof243:  t.cs = 243; goto _test_eof
	_test_eof244:  t.cs = 244; goto _test_eof
	_test_eof245:  t.cs = 245; goto _test_eof
	_test_eof246:  t.cs = 246; goto _test_eof
	_test_eof247:  t.cs = 247; goto _test_eof
	_test_eof248:  t.cs = 248; goto _test_eof
	_test_eof249:  t.cs = 249; goto _test_eof
	_test_eof250:  t.cs = 250; goto _test_eof
	_test_eof251:  t.cs = 251; goto _test_eof
	_test_eof252:  t.cs = 252; goto _test_eof
	_test_eof253:  t.cs = 253; goto _test_eof
	_test_eof683:  t.cs = 683; goto _test_eof
	_test_eof254:  t.cs = 254; goto _test_eof
	_test_eof255:  t.cs = 255; goto _test_eof
	_test_eof256:  t.cs = 256; goto _test_eof
	_test_eof257:  t.cs = 257; goto _test_eof
	_test_eof258:  t.cs = 258; goto _test_eof
	_test_eof259:  t.cs = 259; goto _test_eof
	_test_eof260:  t.cs = 260; goto _test_eof
	_test_eof261:  t.cs = 261; goto _test_eof
	_test_eof262:  t.cs = 262; goto _test_eof
	_test_eof263:  t.cs = 263; goto _test_eof
	_test_eof264:  t.cs = 264; goto _test_eof
	_test_eof265:  t.cs = 265; goto _test_eof
	_test_eof266:  t.cs = 266; goto _test_eof
	_test_eof267:  t.cs = 267; goto _test_eof
	_test_eof268:  t.cs = 268; goto _test_eof
	_test_eof269:  t.cs = 269; goto _test_eof
	_test_eof270:  t.cs = 270; goto _test_eof
	_test_eof271:  t.cs = 271; goto _test_eof
	_test_eof272:  t.cs = 272; goto _test_eof
	_test_eof273:  t.cs = 273; goto _test_eof
	_test_eof274:  t.cs = 274; goto _test_eof
	_test_eof275:  t.cs = 275; goto _test_eof
	_test_eof276:  t.cs = 276; goto _test_eof
	_test_eof277:  t.cs = 277; goto _test_eof
	_test_eof278:  t.cs = 278; goto _test_eof
	_test_eof279:  t.cs = 279; goto _test_eof
	_test_eof280:  t.cs = 280; goto _test_eof
	_test_eof281:  t.cs = 281; goto _test_eof
	_test_eof282:  t.cs = 282; goto _test_eof
	_test_eof283:  t.cs = 283; goto _test_eof
	_test_eof284:  t.cs = 284; goto _test_eof
	_test_eof285:  t.cs = 285; goto _test_eof
	_test_eof286:  t.cs = 286; goto _test_eof
	_test_eof287:  t.cs = 287; goto _test_eof
	_test_eof288:  t.cs = 288; goto _test_eof
	_test_eof289:  t.cs = 289; goto _test_eof
	_test_eof290:  t.cs = 290; goto _test_eof
	_test_eof291:  t.cs = 291; goto _test_eof
	_test_eof292:  t.cs = 292; goto _test_eof
	_test_eof293:  t.cs = 293; goto _test_eof
	_test_eof294:  t.cs = 294; goto _test_eof
	_test_eof295:  t.cs = 295; goto _test_eof
	_test_eof296:  t.cs = 296; goto _test_eof
	_test_eof297:  t.cs = 297; goto _test_eof
	_test_eof298:  t.cs = 298; goto _test_eof
	_test_eof299:  t.cs = 299; goto _test_eof
	_test_eof300:  t.cs = 300; goto _test_eof
	_test_eof301:  t.cs = 301; goto _test_eof
	_test_eof302:  t.cs = 302; goto _test_eof
	_test_eof303:  t.cs = 303; goto _test_eof
	_test_eof304:  t.cs = 304; goto _test_eof
	_test_eof305:  t.cs = 305; goto _test_eof
	_test_eof306:  t.cs = 306; goto _test_eof
	_test_eof307:  t.cs = 307; goto _test_eof
	_test_eof308:  t.cs = 308; goto _test_eof
	_test_eof309:  t.cs = 309; goto _test_eof
	_test_eof310:  t.cs = 310; goto _test_eof
	_test_eof311:  t.cs = 311; goto _test_eof
	_test_eof312:  t.cs = 312; goto _test_eof
	_test_eof313:  t.cs = 313; goto _test_eof
	_test_eof314:  t.cs = 314; goto _test_eof
	_test_eof315:  t.cs = 315; goto _test_eof
	_test_eof316:  t.cs = 316; goto _test_eof
	_test_eof317:  t.cs = 317; goto _test_eof
	_test_eof318:  t.cs = 318; goto _test_eof
	_test_eof319:  t.cs = 319; goto _test_eof
	_test_eof320:  t.cs = 320; goto _test_eof
	_test_eof321:  t.cs = 321; goto _test_eof
	_test_eof322:  t.cs = 322; goto _test_eof
	_test_eof323:  t.cs = 323; goto _test_eof
	_test_eof324:  t.cs = 324; goto _test_eof
	_test_eof325:  t.cs = 325; goto _test_eof
	_test_eof326:  t.cs = 326; goto _test_eof
	_test_eof327:  t.cs = 327; goto _test_eof
	_test_eof328:  t.cs = 328; goto _test_eof
	_test_eof329:  t.cs = 329; goto _test_eof
	_test_eof330:  t.cs = 330; goto _test_eof
	_test_eof331:  t.cs = 331; goto _test_eof
	_test_eof332:  t.cs = 332; goto _test_eof
	_test_eof333:  t.cs = 333; goto _test_eof
	_test_eof334:  t.cs = 334; goto _test_eof
	_test_eof335:  t.cs = 335; goto _test_eof
	_test_eof336:  t.cs = 336; goto _test_eof
	_test_eof337:  t.cs = 337; goto _test_eof
	_test_eof338:  t.cs = 338; goto _test_eof
	_test_eof339:  t.cs = 339; goto _test_eof
	_test_eof340:  t.cs = 340; goto _test_eof
	_test_eof684:  t.cs = 684; goto _test_eof
	_test_eof341:  t.cs = 341; goto _test_eof
	_test_eof342:  t.cs = 342; goto _test_eof
	_test_eof343:  t.cs = 343; goto _test_eof
	_test_eof344:  t.cs = 344; goto _test_eof
	_test_eof345:  t.cs = 345; goto _test_eof
	_test_eof346:  t.cs = 346; goto _test_eof
	_test_eof347:  t.cs = 347; goto _test_eof
	_test_eof348:  t.cs = 348; goto _test_eof
	_test_eof349:  t.cs = 349; goto _test_eof
	_test_eof350:  t.cs = 350; goto _test_eof
	_test_eof351:  t.cs = 351; goto _test_eof
	_test_eof352:  t.cs = 352; goto _test_eof
	_test_eof353:  t.cs = 353; goto _test_eof
	_test_eof354:  t.cs = 354; goto _test_eof
	_test_eof685:  t.cs = 685; goto _test_eof
	_test_eof355:  t.cs = 355; goto _test_eof
	_test_eof356:  t.cs = 356; goto _test_eof
	_test_eof357:  t.cs = 357; goto _test_eof
	_test_eof358:  t.cs = 358; goto _test_eof
	_test_eof359:  t.cs = 359; goto _test_eof
	_test_eof360:  t.cs = 360; goto _test_eof
	_test_eof361:  t.cs = 361; goto _test_eof
	_test_eof362:  t.cs = 362; goto _test_eof
	_test_eof363:  t.cs = 363; goto _test_eof
	_test_eof364:  t.cs = 364; goto _test_eof
	_test_eof365:  t.cs = 365; goto _test_eof
	_test_eof366:  t.cs = 366; goto _test_eof
	_test_eof367:  t.cs = 367; goto _test_eof
	_test_eof368:  t.cs = 368; goto _test_eof
	_test_eof369:  t.cs = 369; goto _test_eof
	_test_eof370:  t.cs = 370; goto _test_eof
	_test_eof371:  t.cs = 371; goto _test_eof
	_test_eof372:  t.cs = 372; goto _test_eof
	_test_eof373:  t.cs = 373; goto _test_eof
	_test_eof374:  t.cs = 374; goto _test_eof
	_test_eof375:  t.cs = 375; goto _test_eof
	_test_eof376:  t.cs = 376; goto _test_eof
	_test_eof377:  t.cs = 377; goto _test_eof
	_test_eof378:  t.cs = 378; goto _test_eof
	_test_eof379:  t.cs = 379; goto _test_eof
	_test_eof380:  t.cs = 380; goto _test_eof
	_test_eof381:  t.cs = 381; goto _test_eof
	_test_eof382:  t.cs = 382; goto _test_eof
	_test_eof383:  t.cs = 383; goto _test_eof
	_test_eof384:  t.cs = 384; goto _test_eof
	_test_eof385:  t.cs = 385; goto _test_eof
	_test_eof386:  t.cs = 386; goto _test_eof
	_test_eof387:  t.cs = 387; goto _test_eof
	_test_eof388:  t.cs = 388; goto _test_eof
	_test_eof389:  t.cs = 389; goto _test_eof
	_test_eof390:  t.cs = 390; goto _test_eof
	_test_eof391:  t.cs = 391; goto _test_eof
	_test_eof392:  t.cs = 392; goto _test_eof
	_test_eof393:  t.cs = 393; goto _test_eof
	_test_eof394:  t.cs = 394; goto _test_eof
	_test_eof395:  t.cs = 395; goto _test_eof
	_test_eof396:  t.cs = 396; goto _test_eof
	_test_eof397:  t.cs = 397; goto _test_eof
	_test_eof398:  t.cs = 398; goto _test_eof
	_test_eof399:  t.cs = 399; goto _test_eof
	_test_eof400:  t.cs = 400; goto _test_eof
	_test_eof401:  t.cs = 401; goto _test_eof
	_test_eof402:  t.cs = 402; goto _test_eof
	_test_eof403:  t.cs = 403; goto _test_eof
	_test_eof404:  t.cs = 404; goto _test_eof
	_test_eof405:  t.cs = 405; goto _test_eof
	_test_eof406:  t.cs = 406; goto _test_eof
	_test_eof407:  t.cs = 407; goto _test_eof
	_test_eof408:  t.cs = 408; goto _test_eof
	_test_eof409:  t.cs = 409; goto _test_eof
	_test_eof410:  t.cs = 410; goto _test_eof
	_test_eof411:  t.cs = 411; goto _test_eof
	_test_eof412:  t.cs = 412; goto _test_eof
	_test_eof413:  t.cs = 413; goto _test_eof
	_test_eof414:  t.cs = 414; goto _test_eof
	_test_eof415:  t.cs = 415; goto _test_eof
	_test_eof416:  t.cs = 416; goto _test_eof
	_test_eof417:  t.cs = 417; goto _test_eof
	_test_eof418:  t.cs = 418; goto _test_eof
	_test_eof419:  t.cs = 419; goto _test_eof
	_test_eof420:  t.cs = 420; goto _test_eof
	_test_eof421:  t.cs = 421; goto _test_eof
	_test_eof422:  t.cs = 422; goto _test_eof
	_test_eof423:  t.cs = 423; goto _test_eof
	_test_eof424:  t.cs = 424; goto _test_eof
	_test_eof425:  t.cs = 425; goto _test_eof
	_test_eof426:  t.cs = 426; goto _test_eof
	_test_eof427:  t.cs = 427; goto _test_eof
	_test_eof428:  t.cs = 428; goto _test_eof
	_test_eof429:  t.cs = 429; goto _test_eof
	_test_eof430:  t.cs = 430; goto _test_eof
	_test_eof431:  t.cs = 431; goto _test_eof
	_test_eof432:  t.cs = 432; goto _test_eof
	_test_eof433:  t.cs = 433; goto _test_eof
	_test_eof434:  t.cs = 434; goto _test_eof
	_test_eof435:  t.cs = 435; goto _test_eof
	_test_eof436:  t.cs = 436; goto _test_eof
	_test_eof437:  t.cs = 437; goto _test_eof
	_test_eof438:  t.cs = 438; goto _test_eof
	_test_eof439:  t.cs = 439; goto _test_eof
	_test_eof440:  t.cs = 440; goto _test_eof
	_test_eof441:  t.cs = 441; goto _test_eof
	_test_eof442:  t.cs = 442; goto _test_eof
	_test_eof443:  t.cs = 443; goto _test_eof
	_test_eof444:  t.cs = 444; goto _test_eof
	_test_eof445:  t.cs = 445; goto _test_eof
	_test_eof446:  t.cs = 446; goto _test_eof
	_test_eof447:  t.cs = 447; goto _test_eof
	_test_eof448:  t.cs = 448; goto _test_eof
	_test_eof449:  t.cs = 449; goto _test_eof
	_test_eof450:  t.cs = 450; goto _test_eof
	_test_eof451:  t.cs = 451; goto _test_eof
	_test_eof452:  t.cs = 452; goto _test_eof
	_test_eof453:  t.cs = 453; goto _test_eof
	_test_eof454:  t.cs = 454; goto _test_eof
	_test_eof455:  t.cs = 455; goto _test_eof
	_test_eof456:  t.cs = 456; goto _test_eof
	_test_eof457:  t.cs = 457; goto _test_eof
	_test_eof458:  t.cs = 458; goto _test_eof
	_test_eof459:  t.cs = 459; goto _test_eof
	_test_eof460:  t.cs = 460; goto _test_eof
	_test_eof461:  t.cs = 461; goto _test_eof
	_test_eof462:  t.cs = 462; goto _test_eof
	_test_eof463:  t.cs = 463; goto _test_eof
	_test_eof464:  t.cs = 464; goto _test_eof
	_test_eof465:  t.cs = 465; goto _test_eof
	_test_eof466:  t.cs = 466; goto _test_eof
	_test_eof467:  t.cs = 467; goto _test_eof
	_test_eof468:  t.cs = 468; goto _test_eof
	_test_eof469:  t.cs = 469; goto _test_eof
	_test_eof470:  t.cs = 470; goto _test_eof
	_test_eof471:  t.cs = 471; goto _test_eof
	_test_eof472:  t.cs = 472; goto _test_eof
	_test_eof473:  t.cs = 473; goto _test_eof
	_test_eof474:  t.cs = 474; goto _test_eof
	_test_eof475:  t.cs = 475; goto _test_eof
	_test_eof476:  t.cs = 476; goto _test_eof
	_test_eof477:  t.cs = 477; goto _test_eof
	_test_eof478:  t.cs = 478; goto _test_eof
	_test_eof479:  t.cs = 479; goto _test_eof
	_test_eof480:  t.cs = 480; goto _test_eof
	_test_eof481:  t.cs = 481; goto _test_eof
	_test_eof482:  t.cs = 482; goto _test_eof
	_test_eof483:  t.cs = 483; goto _test_eof
	_test_eof484:  t.cs = 484; goto _test_eof
	_test_eof485:  t.cs = 485; goto _test_eof
	_test_eof486:  t.cs = 486; goto _test_eof
	_test_eof487:  t.cs = 487; goto _test_eof
	_test_eof488:  t.cs = 488; goto _test_eof
	_test_eof489:  t.cs = 489; goto _test_eof
	_test_eof490:  t.cs = 490; goto _test_eof
	_test_eof686:  t.cs = 686; goto _test_eof
	_test_eof491:  t.cs = 491; goto _test_eof
	_test_eof492:  t.cs = 492; goto _test_eof
	_test_eof687:  t.cs = 687; goto _test_eof
	_test_eof688:  t.cs = 688; goto _test_eof
	_test_eof689:  t.cs = 689; goto _test_eof
	_test_eof690:  t.cs = 690; goto _test_eof
	_test_eof493:  t.cs = 493; goto _test_eof
	_test_eof494:  t.cs = 494; goto _test_eof
	_test_eof495:  t.cs = 495; goto _test_eof
	_test_eof496:  t.cs = 496; goto _test_eof
	_test_eof497:  t.cs = 497; goto _test_eof
	_test_eof498:  t.cs = 498; goto _test_eof
	_test_eof499:  t.cs = 499; goto _test_eof
	_test_eof500:  t.cs = 500; goto _test_eof
	_test_eof501:  t.cs = 501; goto _test_eof
	_test_eof502:  t.cs = 502; goto _test_eof
	_test_eof503:  t.cs = 503; goto _test_eof
	_test_eof504:  t.cs = 504; goto _test_eof
	_test_eof505:  t.cs = 505; goto _test_eof
	_test_eof506:  t.cs = 506; goto _test_eof
	_test_eof507:  t.cs = 507; goto _test_eof
	_test_eof508:  t.cs = 508; goto _test_eof
	_test_eof509:  t.cs = 509; goto _test_eof
	_test_eof510:  t.cs = 510; goto _test_eof
	_test_eof511:  t.cs = 511; goto _test_eof
	_test_eof512:  t.cs = 512; goto _test_eof
	_test_eof513:  t.cs = 513; goto _test_eof
	_test_eof514:  t.cs = 514; goto _test_eof
	_test_eof515:  t.cs = 515; goto _test_eof
	_test_eof516:  t.cs = 516; goto _test_eof
	_test_eof517:  t.cs = 517; goto _test_eof
	_test_eof518:  t.cs = 518; goto _test_eof
	_test_eof519:  t.cs = 519; goto _test_eof
	_test_eof520:  t.cs = 520; goto _test_eof
	_test_eof521:  t.cs = 521; goto _test_eof
	_test_eof522:  t.cs = 522; goto _test_eof
	_test_eof523:  t.cs = 523; goto _test_eof
	_test_eof524:  t.cs = 524; goto _test_eof
	_test_eof525:  t.cs = 525; goto _test_eof
	_test_eof526:  t.cs = 526; goto _test_eof
	_test_eof527:  t.cs = 527; goto _test_eof
	_test_eof528:  t.cs = 528; goto _test_eof
	_test_eof529:  t.cs = 529; goto _test_eof
	_test_eof530:  t.cs = 530; goto _test_eof
	_test_eof531:  t.cs = 531; goto _test_eof
	_test_eof532:  t.cs = 532; goto _test_eof
	_test_eof691:  t.cs = 691; goto _test_eof
	_test_eof533:  t.cs = 533; goto _test_eof
	_test_eof534:  t.cs = 534; goto _test_eof
	_test_eof535:  t.cs = 535; goto _test_eof
	_test_eof536:  t.cs = 536; goto _test_eof
	_test_eof537:  t.cs = 537; goto _test_eof
	_test_eof538:  t.cs = 538; goto _test_eof
	_test_eof539:  t.cs = 539; goto _test_eof
	_test_eof540:  t.cs = 540; goto _test_eof
	_test_eof541:  t.cs = 541; goto _test_eof
	_test_eof542:  t.cs = 542; goto _test_eof
	_test_eof543:  t.cs = 543; goto _test_eof
	_test_eof544:  t.cs = 544; goto _test_eof
	_test_eof545:  t.cs = 545; goto _test_eof
	_test_eof546:  t.cs = 546; goto _test_eof
	_test_eof692:  t.cs = 692; goto _test_eof
	_test_eof547:  t.cs = 547; goto _test_eof
	_test_eof548:  t.cs = 548; goto _test_eof
	_test_eof549:  t.cs = 549; goto _test_eof
	_test_eof550:  t.cs = 550; goto _test_eof
	_test_eof551:  t.cs = 551; goto _test_eof
	_test_eof552:  t.cs = 552; goto _test_eof
	_test_eof553:  t.cs = 553; goto _test_eof
	_test_eof554:  t.cs = 554; goto _test_eof
	_test_eof555:  t.cs = 555; goto _test_eof
	_test_eof556:  t.cs = 556; goto _test_eof
	_test_eof557:  t.cs = 557; goto _test_eof
	_test_eof558:  t.cs = 558; goto _test_eof
	_test_eof559:  t.cs = 559; goto _test_eof
	_test_eof560:  t.cs = 560; goto _test_eof
	_test_eof561:  t.cs = 561; goto _test_eof
	_test_eof562:  t.cs = 562; goto _test_eof
	_test_eof563:  t.cs = 563; goto _test_eof
	_test_eof564:  t.cs = 564; goto _test_eof
	_test_eof565:  t.cs = 565; goto _test_eof
	_test_eof566:  t.cs = 566; goto _test_eof
	_test_eof693:  t.cs = 693; goto _test_eof
	_test_eof694:  t.cs = 694; goto _test_eof
	_test_eof567:  t.cs = 567; goto _test_eof
	_test_eof568:  t.cs = 568; goto _test_eof
	_test_eof569:  t.cs = 569; goto _test_eof
	_test_eof570:  t.cs = 570; goto _test_eof
	_test_eof571:  t.cs = 571; goto _test_eof
	_test_eof695:  t.cs = 695; goto _test_eof
	_test_eof696:  t.cs = 696; goto _test_eof
	_test_eof572:  t.cs = 572; goto _test_eof
	_test_eof573:  t.cs = 573; goto _test_eof
	_test_eof574:  t.cs = 574; goto _test_eof
	_test_eof575:  t.cs = 575; goto _test_eof
	_test_eof576:  t.cs = 576; goto _test_eof
	_test_eof577:  t.cs = 577; goto _test_eof
	_test_eof578:  t.cs = 578; goto _test_eof
	_test_eof579:  t.cs = 579; goto _test_eof
	_test_eof580:  t.cs = 580; goto _test_eof
	_test_eof581:  t.cs = 581; goto _test_eof
	_test_eof582:  t.cs = 582; goto _test_eof
	_test_eof583:  t.cs = 583; goto _test_eof
	_test_eof584:  t.cs = 584; goto _test_eof
	_test_eof585:  t.cs = 585; goto _test_eof
	_test_eof586:  t.cs = 586; goto _test_eof
	_test_eof587:  t.cs = 587; goto _test_eof
	_test_eof588:  t.cs = 588; goto _test_eof
	_test_eof589:  t.cs = 589; goto _test_eof
	_test_eof590:  t.cs = 590; goto _test_eof
	_test_eof591:  t.cs = 591; goto _test_eof
	_test_eof592:  t.cs = 592; goto _test_eof
	_test_eof593:  t.cs = 593; goto _test_eof
	_test_eof594:  t.cs = 594; goto _test_eof
	_test_eof595:  t.cs = 595; goto _test_eof
	_test_eof596:  t.cs = 596; goto _test_eof
	_test_eof597:  t.cs = 597; goto _test_eof
	_test_eof598:  t.cs = 598; goto _test_eof
	_test_eof599:  t.cs = 599; goto _test_eof
	_test_eof600:  t.cs = 600; goto _test_eof
	_test_eof601:  t.cs = 601; goto _test_eof
	_test_eof602:  t.cs = 602; goto _test_eof
	_test_eof603:  t.cs = 603; goto _test_eof
	_test_eof604:  t.cs = 604; goto _test_eof
	_test_eof605:  t.cs = 605; goto _test_eof
	_test_eof606:  t.cs = 606; goto _test_eof
	_test_eof607:  t.cs = 607; goto _test_eof
	_test_eof608:  t.cs = 608; goto _test_eof
	_test_eof609:  t.cs = 609; goto _test_eof
	_test_eof610:  t.cs = 610; goto _test_eof
	_test_eof697:  t.cs = 697; goto _test_eof
	_test_eof611:  t.cs = 611; goto _test_eof
	_test_eof612:  t.cs = 612; goto _test_eof
	_test_eof613:  t.cs = 613; goto _test_eof
	_test_eof614:  t.cs = 614; goto _test_eof
	_test_eof615:  t.cs = 615; goto _test_eof
	_test_eof616:  t.cs = 616; goto _test_eof
	_test_eof617:  t.cs = 617; goto _test_eof
	_test_eof618:  t.cs = 618; goto _test_eof
	_test_eof619:  t.cs = 619; goto _test_eof
	_test_eof620:  t.cs = 620; goto _test_eof
	_test_eof621:  t.cs = 621; goto _test_eof
	_test_eof622:  t.cs = 622; goto _test_eof
	_test_eof623:  t.cs = 623; goto _test_eof
	_test_eof624:  t.cs = 624; goto _test_eof
	_test_eof625:  t.cs = 625; goto _test_eof
	_test_eof626:  t.cs = 626; goto _test_eof
	_test_eof627:  t.cs = 627; goto _test_eof
	_test_eof628:  t.cs = 628; goto _test_eof
	_test_eof629:  t.cs = 629; goto _test_eof
	_test_eof630:  t.cs = 630; goto _test_eof
	_test_eof631:  t.cs = 631; goto _test_eof
	_test_eof632:  t.cs = 632; goto _test_eof
	_test_eof633:  t.cs = 633; goto _test_eof
	_test_eof634:  t.cs = 634; goto _test_eof
	_test_eof635:  t.cs = 635; goto _test_eof
	_test_eof636:  t.cs = 636; goto _test_eof
	_test_eof637:  t.cs = 637; goto _test_eof
	_test_eof638:  t.cs = 638; goto _test_eof
	_test_eof639:  t.cs = 639; goto _test_eof
	_test_eof640:  t.cs = 640; goto _test_eof
	_test_eof641:  t.cs = 641; goto _test_eof
	_test_eof642:  t.cs = 642; goto _test_eof
	_test_eof643:  t.cs = 643; goto _test_eof
	_test_eof644:  t.cs = 644; goto _test_eof
	_test_eof645:  t.cs = 645; goto _test_eof
	_test_eof646:  t.cs = 646; goto _test_eof
	_test_eof647:  t.cs = 647; goto _test_eof
	_test_eof648:  t.cs = 648; goto _test_eof
	_test_eof649:  t.cs = 649; goto _test_eof
	_test_eof650:  t.cs = 650; goto _test_eof
	_test_eof651:  t.cs = 651; goto _test_eof
	_test_eof652:  t.cs = 652; goto _test_eof
	_test_eof653:  t.cs = 653; goto _test_eof
	_test_eof654:  t.cs = 654; goto _test_eof
	_test_eof655:  t.cs = 655; goto _test_eof
	_test_eof656:  t.cs = 656; goto _test_eof
	_test_eof657:  t.cs = 657; goto _test_eof
	_test_eof658:  t.cs = 658; goto _test_eof
	_test_eof659:  t.cs = 659; goto _test_eof
	_test_eof660:  t.cs = 660; goto _test_eof
	_test_eof661:  t.cs = 661; goto _test_eof
	_test_eof662:  t.cs = 662; goto _test_eof
	_test_eof663:  t.cs = 663; goto _test_eof
	_test_eof664:  t.cs = 664; goto _test_eof
	_test_eof665:  t.cs = 665; goto _test_eof
	_test_eof666:  t.cs = 666; goto _test_eof
	_test_eof667:  t.cs = 667; goto _test_eof
	_test_eof668:  t.cs = 668; goto _test_eof
	_test_eof669:  t.cs = 669; goto _test_eof
	_test_eof670:  t.cs = 670; goto _test_eof
	_test_eof671:  t.cs = 671; goto _test_eof
	_test_eof672:  t.cs = 672; goto _test_eof
	_test_eof673:  t.cs = 673; goto _test_eof
	_test_eof674:  t.cs = 674; goto _test_eof
	_test_eof675:  t.cs = 675; goto _test_eof
	_test_eof676:  t.cs = 676; goto _test_eof
	_test_eof677:  t.cs = 677; goto _test_eof

	_test_eof: {}
	if ( t.p) == ( t.pe) {
		switch  t.cs {
		case 681, 682, 684, 691, 692:
//line tokenizer.go.rl:76

		t.tokval(StrTok(t.data[t.prev:t.p]))
	
		case 678, 679:
//line tokenizer.go.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 0; goto _out }
		}
		t.tokval(IntTok(temp.Int))
	
		case 688:
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
		case 686, 687:
//line tokenizer.go.rl:140
t.tok(SOURCE)
		case 689, 690:
//line tokenizer.go.rl:135
t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))
//line tokenizer.go.rl:140
t.tok(SOURCE)
//line tokenizer.go:12244
		}
	}

	_out: {}
	}

//line tokenizer.go.rl:228

	if debugTokenizer {
		fmt.Printf("exited: %s\n", t.state)
	}
	if t.p != t.pe {
		t.err(fmt.Errorf("%w: %q %q", ErrLineIsNotFinished, t.data[:t.p], t.data[t.p:t.pe]))
	}
	return t.tokens, errors.Join(t.errors...)
}

func (t *tokenizer) parseTime(s string) (time.Time, error) {
	return ParseTime(t.nowTime, s)
}

func (t *tokenizer) tokval(token Token) {
	t.tokens = append(t.tokens, token)
}

func (t *tokenizer) tok(tok int) {
	t.tokens = append(t.tokens, VoidTok(tok))
}

func (t *tokenizer) err(err error) {
	t.errors = append(t.errors, err)
}