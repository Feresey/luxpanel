
//line ./combat/tokenizer.go.rl:1
// Generated code. DO NOT EDIT

package combat

import (
	"fmt"
	"strconv"
	"time"
	"errors"
)


//line ./combat/tokenizer.go:14
const logparser_start int = 1
const logparser_first_final int = 522
const logparser_error int = 0

const logparser_en_main int = 1


//line ./combat/tokenizer.go.rl:20


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
	tokens []token
	errors []error
	state
}


//line ./combat/tokenizer.go.rl:184


func (t *tokenizer) Parse(nowTime time.Time, data string) ([]token, error) {
	var parseErr error
	var temp yySymType
	var Float float64

	t.state = state{}
	
//line ./combat/tokenizer.go:79
	{
	 t.cs = logparser_start
	}

//line ./combat/tokenizer.go.rl:193

	t.prev = 0
	t.pe = len(data)
	t.nowTime = nowTime
	t.data = data

	t.tokens = t.tokens[:0]
	t.errors = nil

	
//line ./combat/tokenizer.go:93
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
	case 522:
		goto st_case_522
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
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
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
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
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
	case 527:
		goto st_case_527
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
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
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
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
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
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
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
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
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
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
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
	case 538:
		goto st_case_538
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
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
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
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:1206
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
//line ./combat/tokenizer.go.rl:98

		temp.Time, parseErr = parseTime(t.nowTime, t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: TIME, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("parseTime: %w", parseErr)})
			{( t.p)++;  t.cs = 14; goto _out }
		}
		t.tokval(timeTok(temp.Time))
	
	goto st14
	st14:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof14
		}
	st_case_14:
//line ./combat/tokenizer.go:1326
		switch ( t.data)[( t.p)] {
		case 32:
			goto st14
		case 67:
			goto tr15
		}
		goto st0
tr15:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:1348
		if ( t.data)[( t.p)] == 77 {
			goto st16
		}
		goto st0
	st16:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof16
		}
	st_case_16:
		if ( t.data)[( t.p)] == 66 {
			goto st17
		}
		goto st0
	st17:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof17
		}
	st_case_17:
		if ( t.data)[( t.p)] == 84 {
			goto st18
		}
		goto st0
	st18:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof18
		}
	st_case_18:
		if ( t.data)[( t.p)] == 32 {
			goto st19
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st19
		}
		goto st0
	st19:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st19
		case 124:
			goto tr20
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st19
		}
		goto st0
tr20:
//line ./combat/tokenizer.go.rl:122
 t.tok(COMBAT) 
	goto st20
	st20:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof20
		}
	st_case_20:
//line ./combat/tokenizer.go:1407
		if ( t.data)[( t.p)] == 32 {
			goto st21
		}
		goto st0
	st21:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st21
		case 61:
			goto tr22
		case 68:
			goto tr23
		case 71:
			goto tr24
		case 72:
			goto tr25
		case 75:
			goto tr26
		case 80:
			goto tr27
		case 82:
			goto tr28
		}
		goto st0
tr22:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st22
	st22:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof22
		}
	st_case_22:
//line ./combat/tokenizer.go:1450
		if ( t.data)[( t.p)] == 61 {
			goto st23
		}
		goto st0
	st23:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof23
		}
	st_case_23:
		if ( t.data)[( t.p)] == 61 {
			goto st24
		}
		goto st0
	st24:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof24
		}
	st_case_24:
		if ( t.data)[( t.p)] == 61 {
			goto st25
		}
		goto st0
	st25:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof25
		}
	st_case_25:
		if ( t.data)[( t.p)] == 61 {
			goto st26
		}
		goto st0
	st26:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof26
		}
	st_case_26:
		if ( t.data)[( t.p)] == 61 {
			goto st27
		}
		goto st0
	st27:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof27
		}
	st_case_27:
		if ( t.data)[( t.p)] == 61 {
			goto st28
		}
		goto st0
	st28:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof28
		}
	st_case_28:
		if ( t.data)[( t.p)] == 32 {
			goto st29
		}
		goto st0
	st29:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch ( t.data)[( t.p)] {
		case 67:
			goto st30
		case 83:
			goto st63
		}
		goto st0
	st30:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof30
		}
	st_case_30:
		if ( t.data)[( t.p)] == 111 {
			goto st31
		}
		goto st0
	st31:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof31
		}
	st_case_31:
		if ( t.data)[( t.p)] == 110 {
			goto st32
		}
		goto st0
	st32:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof32
		}
	st_case_32:
		if ( t.data)[( t.p)] == 110 {
			goto st33
		}
		goto st0
	st33:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof33
		}
	st_case_33:
		if ( t.data)[( t.p)] == 101 {
			goto st34
		}
		goto st0
	st34:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof34
		}
	st_case_34:
		if ( t.data)[( t.p)] == 99 {
			goto st35
		}
		goto st0
	st35:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof35
		}
	st_case_35:
		if ( t.data)[( t.p)] == 116 {
			goto st36
		}
		goto st0
	st36:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof36
		}
	st_case_36:
		if ( t.data)[( t.p)] == 32 {
			goto st37
		}
		goto st0
	st37:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof37
		}
	st_case_37:
		if ( t.data)[( t.p)] == 116 {
			goto st38
		}
		goto st0
	st38:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof38
		}
	st_case_38:
		if ( t.data)[( t.p)] == 111 {
			goto st39
		}
		goto st0
	st39:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof39
		}
	st_case_39:
		if ( t.data)[( t.p)] == 32 {
			goto st40
		}
		goto st0
	st40:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof40
		}
	st_case_40:
		if ( t.data)[( t.p)] == 103 {
			goto st41
		}
		goto st0
	st41:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof41
		}
	st_case_41:
		if ( t.data)[( t.p)] == 97 {
			goto st42
		}
		goto st0
	st42:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof42
		}
	st_case_42:
		if ( t.data)[( t.p)] == 109 {
			goto st43
		}
		goto st0
	st43:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof43
		}
	st_case_43:
		if ( t.data)[( t.p)] == 101 {
			goto st44
		}
		goto st0
	st44:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof44
		}
	st_case_44:
		if ( t.data)[( t.p)] == 32 {
			goto st45
		}
		goto st0
	st45:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof45
		}
	st_case_45:
		if ( t.data)[( t.p)] == 115 {
			goto st46
		}
		goto st0
	st46:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof46
		}
	st_case_46:
		if ( t.data)[( t.p)] == 101 {
			goto st47
		}
		goto st0
	st47:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof47
		}
	st_case_47:
		if ( t.data)[( t.p)] == 115 {
			goto st48
		}
		goto st0
	st48:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof48
		}
	st_case_48:
		if ( t.data)[( t.p)] == 115 {
			goto st49
		}
		goto st0
	st49:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof49
		}
	st_case_49:
		if ( t.data)[( t.p)] == 105 {
			goto st50
		}
		goto st0
	st50:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof50
		}
	st_case_50:
		if ( t.data)[( t.p)] == 111 {
			goto st51
		}
		goto st0
	st51:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof51
		}
	st_case_51:
		if ( t.data)[( t.p)] == 110 {
			goto st52
		}
		goto st0
	st52:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof52
		}
	st_case_52:
		if ( t.data)[( t.p)] == 32 {
			goto st53
		}
		goto st0
	st53:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof53
		}
	st_case_53:
		if ( t.data)[( t.p)] == 45 {
			goto tr61
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr62
		}
		goto st0
tr61:
//line ./combat/tokenizer.go.rl:168
t.tok(CONNECT_TO_GAME_SESSION_PREFIX)
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st54
	st54:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof54
		}
	st_case_54:
//line ./combat/tokenizer.go:1756
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st55
		}
		goto st0
tr62:
//line ./combat/tokenizer.go.rl:168
t.tok(CONNECT_TO_GAME_SESSION_PREFIX)
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st55
	st55:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof55
		}
	st_case_55:
//line ./combat/tokenizer.go:1777
		if ( t.data)[( t.p)] == 32 {
			goto tr64
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st55
		}
		goto st0
tr64:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 56; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st56
	st56:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof56
		}
	st_case_56:
//line ./combat/tokenizer.go:1801
		if ( t.data)[( t.p)] == 61 {
			goto st57
		}
		goto st0
tr132:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 57; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st57
	st57:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof57
		}
	st_case_57:
//line ./combat/tokenizer.go:1822
		if ( t.data)[( t.p)] == 61 {
			goto st58
		}
		goto st0
	st58:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof58
		}
	st_case_58:
		if ( t.data)[( t.p)] == 61 {
			goto st59
		}
		goto st0
	st59:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof59
		}
	st_case_59:
		if ( t.data)[( t.p)] == 61 {
			goto st60
		}
		goto st0
	st60:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof60
		}
	st_case_60:
		if ( t.data)[( t.p)] == 61 {
			goto st61
		}
		goto st0
	st61:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof61
		}
	st_case_61:
		if ( t.data)[( t.p)] == 61 {
			goto st62
		}
		goto st0
	st62:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof62
		}
	st_case_62:
		if ( t.data)[( t.p)] == 61 {
			goto st522
		}
		goto st0
tr599:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 522; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st522
tr442:
//line ./combat/tokenizer.go.rl:146
t.tok(FRIENDLY_FIRE)
	goto st522
tr600:
//line ./combat/tokenizer.go.rl:137
t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))
	goto st522
tr611:
//line ./combat/tokenizer.go.rl:149
t.tok(PARTICIPATION_MODIFIERS_END)
	goto st522
tr616:
//line ./combat/tokenizer.go.rl:171
t.tokval(strTok(t.data[t.p:]))
	goto st522
	st522:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof522
		}
	st_case_522:
//line ./combat/tokenizer.go:1904
		if ( t.data)[( t.p)] == 32 {
			goto st522
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
	st63:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof63
		}
	st_case_63:
		if ( t.data)[( t.p)] == 116 {
			goto st64
		}
		goto st0
	st64:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof64
		}
	st_case_64:
		if ( t.data)[( t.p)] == 97 {
			goto st65
		}
		goto st0
	st65:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof65
		}
	st_case_65:
		if ( t.data)[( t.p)] == 114 {
			goto st66
		}
		goto st0
	st66:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof66
		}
	st_case_66:
		if ( t.data)[( t.p)] == 116 {
			goto tr75
		}
		goto st0
tr75:
//line ./combat/tokenizer.go.rl:129
 t.tok(START) 
	goto st67
	st67:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof67
		}
	st_case_67:
//line ./combat/tokenizer.go:1957
		if ( t.data)[( t.p)] == 32 {
			goto st68
		}
		goto st0
	st68:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch ( t.data)[( t.p)] {
		case 80:
			goto tr77
		case 103:
			goto tr78
		}
		goto st0
tr77:
//line ./combat/tokenizer.go.rl:164
t.tokval(strTok("PVE mission"))
	goto st69
	st69:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof69
		}
	st_case_69:
//line ./combat/tokenizer.go:1983
		if ( t.data)[( t.p)] == 86 {
			goto st70
		}
		goto st0
	st70:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof70
		}
	st_case_70:
		if ( t.data)[( t.p)] == 69 {
			goto st71
		}
		goto st0
	st71:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof71
		}
	st_case_71:
		if ( t.data)[( t.p)] == 32 {
			goto st72
		}
		goto st0
	st72:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof72
		}
	st_case_72:
		if ( t.data)[( t.p)] == 109 {
			goto st73
		}
		goto st0
	st73:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof73
		}
	st_case_73:
		if ( t.data)[( t.p)] == 105 {
			goto st74
		}
		goto st0
	st74:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof74
		}
	st_case_74:
		if ( t.data)[( t.p)] == 115 {
			goto st75
		}
		goto st0
	st75:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof75
		}
	st_case_75:
		if ( t.data)[( t.p)] == 115 {
			goto st76
		}
		goto st0
	st76:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof76
		}
	st_case_76:
		if ( t.data)[( t.p)] == 105 {
			goto st77
		}
		goto st0
	st77:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof77
		}
	st_case_77:
		if ( t.data)[( t.p)] == 111 {
			goto st78
		}
		goto st0
	st78:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof78
		}
	st_case_78:
		if ( t.data)[( t.p)] == 110 {
			goto st79
		}
		goto st0
	st79:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof79
		}
	st_case_79:
		if ( t.data)[( t.p)] == 32 {
			goto st80
		}
		goto st0
	st80:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof80
		}
	st_case_80:
		if ( t.data)[( t.p)] == 39 {
			goto st81
		}
		goto st0
	st81:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr91
		case 95:
			goto tr92
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr92
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr92
			}
		default:
			goto tr92
		}
		goto st0
tr91:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st82
	st82:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof82
		}
	st_case_82:
//line ./combat/tokenizer.go:2125
		if ( t.data)[( t.p)] == 95 {
			goto st83
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st83
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st83
			}
		default:
			goto st83
		}
		goto st0
	st83:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st84
		case 95:
			goto st83
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st83
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st83
			}
		default:
			goto st83
		}
		goto st0
	st84:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof84
		}
	st_case_84:
		if ( t.data)[( t.p)] == 39 {
			goto tr95
		}
		goto st0
tr95:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st85
	st85:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof85
		}
	st_case_85:
//line ./combat/tokenizer.go:2189
		if ( t.data)[( t.p)] == 32 {
			goto st86
		}
		goto st0
	st86:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof86
		}
	st_case_86:
		if ( t.data)[( t.p)] == 109 {
			goto st87
		}
		goto st0
	st87:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof87
		}
	st_case_87:
		if ( t.data)[( t.p)] == 97 {
			goto st88
		}
		goto st0
	st88:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof88
		}
	st_case_88:
		if ( t.data)[( t.p)] == 112 {
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
		if ( t.data)[( t.p)] == 39 {
			goto st91
		}
		goto st0
	st91:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr102
		case 95:
			goto tr103
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr103
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr103
			}
		default:
			goto tr103
		}
		goto st0
tr102:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st92
	st92:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof92
		}
	st_case_92:
//line ./combat/tokenizer.go:2277
		if ( t.data)[( t.p)] == 95 {
			goto st93
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st93
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st93
			}
		default:
			goto st93
		}
		goto st0
	st93:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st94
		case 95:
			goto st93
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st93
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st93
			}
		default:
			goto st93
		}
		goto st0
	st94:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof94
		}
	st_case_94:
		if ( t.data)[( t.p)] == 39 {
			goto tr106
		}
		goto st0
tr106:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st95
	st95:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof95
		}
	st_case_95:
//line ./combat/tokenizer.go:2341
		switch ( t.data)[( t.p)] {
		case 32:
			goto st96
		case 44:
			goto st97
		case 61:
			goto st57
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st96
		}
		goto st0
tr131:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 96; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st96
	st96:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof96
		}
	st_case_96:
//line ./combat/tokenizer.go:2370
		switch ( t.data)[( t.p)] {
		case 32:
			goto st96
		case 61:
			goto st57
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st96
		}
		goto st0
	st97:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof97
		}
	st_case_97:
		if ( t.data)[( t.p)] == 32 {
			goto st98
		}
		goto st0
	st98:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof98
		}
	st_case_98:
		if ( t.data)[( t.p)] == 108 {
			goto st99
		}
		goto st0
	st99:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof99
		}
	st_case_99:
		if ( t.data)[( t.p)] == 111 {
			goto st100
		}
		goto st0
	st100:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof100
		}
	st_case_100:
		if ( t.data)[( t.p)] == 99 {
			goto st101
		}
		goto st0
	st101:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof101
		}
	st_case_101:
		if ( t.data)[( t.p)] == 97 {
			goto st102
		}
		goto st0
	st102:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof102
		}
	st_case_102:
		if ( t.data)[( t.p)] == 108 {
			goto st103
		}
		goto st0
	st103:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof103
		}
	st_case_103:
		if ( t.data)[( t.p)] == 32 {
			goto st104
		}
		goto st0
	st104:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof104
		}
	st_case_104:
		if ( t.data)[( t.p)] == 99 {
			goto st105
		}
		goto st0
	st105:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof105
		}
	st_case_105:
		if ( t.data)[( t.p)] == 108 {
			goto st106
		}
		goto st0
	st106:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof106
		}
	st_case_106:
		if ( t.data)[( t.p)] == 105 {
			goto st107
		}
		goto st0
	st107:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof107
		}
	st_case_107:
		if ( t.data)[( t.p)] == 101 {
			goto st108
		}
		goto st0
	st108:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof108
		}
	st_case_108:
		if ( t.data)[( t.p)] == 110 {
			goto st109
		}
		goto st0
	st109:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof109
		}
	st_case_109:
		if ( t.data)[( t.p)] == 116 {
			goto st110
		}
		goto st0
	st110:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof110
		}
	st_case_110:
		if ( t.data)[( t.p)] == 32 {
			goto st111
		}
		goto st0
	st111:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof111
		}
	st_case_111:
		if ( t.data)[( t.p)] == 116 {
			goto st112
		}
		goto st0
	st112:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof112
		}
	st_case_112:
		if ( t.data)[( t.p)] == 101 {
			goto st113
		}
		goto st0
	st113:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof113
		}
	st_case_113:
		if ( t.data)[( t.p)] == 97 {
			goto st114
		}
		goto st0
	st114:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof114
		}
	st_case_114:
		if ( t.data)[( t.p)] == 109 {
			goto st115
		}
		goto st0
	st115:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof115
		}
	st_case_115:
		if ( t.data)[( t.p)] == 32 {
			goto st116
		}
		goto st0
	st116:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof116
		}
	st_case_116:
		if ( t.data)[( t.p)] == 45 {
			goto tr128
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr129
		}
		goto st0
tr128:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st117
	st117:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof117
		}
	st_case_117:
//line ./combat/tokenizer.go:2578
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st118
		}
		goto st0
tr129:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st118
	st118:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof118
		}
	st_case_118:
//line ./combat/tokenizer.go:2597
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr131
		case 61:
			goto tr132
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st118
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr131
		}
		goto st0
tr103:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st119
	st119:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof119
		}
	st_case_119:
//line ./combat/tokenizer.go:2627
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr106
		case 95:
			goto st119
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st119
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st119
			}
		default:
			goto st119
		}
		goto st0
tr92:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st120
	st120:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof120
		}
	st_case_120:
//line ./combat/tokenizer.go:2661
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr95
		case 95:
			goto st120
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st120
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
tr78:
//line ./combat/tokenizer.go.rl:163
t.tokval(strTok("gameplay"))
	goto st121
	st121:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof121
		}
	st_case_121:
//line ./combat/tokenizer.go:2690
		if ( t.data)[( t.p)] == 97 {
			goto st122
		}
		goto st0
	st122:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof122
		}
	st_case_122:
		if ( t.data)[( t.p)] == 109 {
			goto st123
		}
		goto st0
	st123:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof123
		}
	st_case_123:
		if ( t.data)[( t.p)] == 101 {
			goto st124
		}
		goto st0
	st124:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof124
		}
	st_case_124:
		if ( t.data)[( t.p)] == 112 {
			goto st125
		}
		goto st0
	st125:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof125
		}
	st_case_125:
		if ( t.data)[( t.p)] == 108 {
			goto st126
		}
		goto st0
	st126:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof126
		}
	st_case_126:
		if ( t.data)[( t.p)] == 97 {
			goto st127
		}
		goto st0
	st127:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof127
		}
	st_case_127:
		if ( t.data)[( t.p)] == 121 {
			goto st79
		}
		goto st0
tr23:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st128
	st128:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof128
		}
	st_case_128:
//line ./combat/tokenizer.go:2763
		if ( t.data)[( t.p)] == 97 {
			goto st129
		}
		goto st0
	st129:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof129
		}
	st_case_129:
		if ( t.data)[( t.p)] == 109 {
			goto st130
		}
		goto st0
	st130:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof130
		}
	st_case_130:
		if ( t.data)[( t.p)] == 97 {
			goto st131
		}
		goto st0
	st131:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof131
		}
	st_case_131:
		if ( t.data)[( t.p)] == 103 {
			goto st132
		}
		goto st0
	st132:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof132
		}
	st_case_132:
		if ( t.data)[( t.p)] == 101 {
			goto tr145
		}
		goto st0
tr145:
//line ./combat/tokenizer.go.rl:125
 t.tok(DAMAGE) 
	goto st133
	st133:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof133
		}
	st_case_133:
//line ./combat/tokenizer.go:2813
		if ( t.data)[( t.p)] == 32 {
			goto st134
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st134
		}
		goto st0
	st134:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st134
		case 40:
			goto tr147
		case 95:
			goto tr148
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st134
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr148
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr148
			}
		default:
			goto tr148
		}
		goto st0
tr147:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st135
	st135:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof135
		}
	st_case_135:
//line ./combat/tokenizer.go:2866
		if ( t.data)[( t.p)] == 95 {
			goto st136
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st136
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st136
			}
		default:
			goto st136
		}
		goto st0
	st136:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st137
		case 95:
			goto st136
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st136
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st136
			}
		default:
			goto st136
		}
		goto st0
	st137:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr151
		case 124:
			goto tr152
		}
		goto st0
tr151:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st138
	st138:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof138
		}
	st_case_138:
//line ./combat/tokenizer.go:2935
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr153
		case 95:
			goto tr154
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr154
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr154
			}
		default:
			goto tr154
		}
		goto st0
tr153:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st139
	st139:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof139
		}
	st_case_139:
//line ./combat/tokenizer.go:2969
		if ( t.data)[( t.p)] == 95 {
			goto st140
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st140
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st140
			}
		default:
			goto st140
		}
		goto st0
	st140:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof140
		}
	st_case_140:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st141
		case 95:
			goto st140
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st140
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st140
			}
		default:
			goto st140
		}
		goto st0
	st141:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof141
		}
	st_case_141:
		if ( t.data)[( t.p)] == 41 {
			goto tr157
		}
		goto st0
tr157:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st142
	st142:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof142
		}
	st_case_142:
//line ./combat/tokenizer.go:3035
		if ( t.data)[( t.p)] == 124 {
			goto tr158
		}
		goto st0
tr152:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st143
tr158:
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st143
	st143:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof143
		}
	st_case_143:
//line ./combat/tokenizer.go:3060
		if ( t.data)[( t.p)] == 45 {
			goto tr159
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr160
		}
		goto st0
tr159:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st144
	st144:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof144
		}
	st_case_144:
//line ./combat/tokenizer.go:3082
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st145
		}
		goto st0
tr160:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st145
	st145:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof145
		}
	st_case_145:
//line ./combat/tokenizer.go:3101
		if ( t.data)[( t.p)] == 32 {
			goto tr162
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st145
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr162
		}
		goto st0
tr162:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 146; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st146
	st146:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof146
		}
	st_case_146:
//line ./combat/tokenizer.go:3130
		switch ( t.data)[( t.p)] {
		case 32:
			goto st146
		case 45:
			goto st147
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st146
		}
		goto st0
	st147:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof147
		}
	st_case_147:
		if ( t.data)[( t.p)] == 62 {
			goto tr165
		}
		goto st0
tr165:
//line ./combat/tokenizer.go.rl:143
t.tok(ARROW)
	goto st148
	st148:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof148
		}
	st_case_148:
//line ./combat/tokenizer.go:3159
		if ( t.data)[( t.p)] == 32 {
			goto st149
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st149
		}
		goto st0
	st149:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st149
		case 40:
			goto tr167
		case 95:
			goto tr168
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st149
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr168
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr168
			}
		default:
			goto tr168
		}
		goto st0
tr167:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st150
	st150:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof150
		}
	st_case_150:
//line ./combat/tokenizer.go:3212
		if ( t.data)[( t.p)] == 95 {
			goto st151
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st151
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st151
			}
		default:
			goto st151
		}
		goto st0
	st151:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof151
		}
	st_case_151:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st152
		case 95:
			goto st151
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st151
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st151
			}
		default:
			goto st151
		}
		goto st0
	st152:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof152
		}
	st_case_152:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr171
		case 124:
			goto tr172
		}
		goto st0
tr171:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st153
	st153:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof153
		}
	st_case_153:
//line ./combat/tokenizer.go:3281
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr173
		case 95:
			goto tr174
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr174
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
tr173:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st154
	st154:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof154
		}
	st_case_154:
//line ./combat/tokenizer.go:3315
		if ( t.data)[( t.p)] == 95 {
			goto st155
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st155
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st155
			}
		default:
			goto st155
		}
		goto st0
	st155:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof155
		}
	st_case_155:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st156
		case 95:
			goto st155
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st155
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st155
			}
		default:
			goto st155
		}
		goto st0
	st156:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof156
		}
	st_case_156:
		if ( t.data)[( t.p)] == 41 {
			goto tr177
		}
		goto st0
tr177:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st157
	st157:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof157
		}
	st_case_157:
//line ./combat/tokenizer.go:3381
		if ( t.data)[( t.p)] == 124 {
			goto tr178
		}
		goto st0
tr172:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st158
tr178:
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st158
	st158:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof158
		}
	st_case_158:
//line ./combat/tokenizer.go:3406
		if ( t.data)[( t.p)] == 45 {
			goto tr179
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr180
		}
		goto st0
tr179:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st159
	st159:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof159
		}
	st_case_159:
//line ./combat/tokenizer.go:3428
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st160
		}
		goto st0
tr180:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st160
	st160:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof160
		}
	st_case_160:
//line ./combat/tokenizer.go:3447
		if ( t.data)[( t.p)] == 32 {
			goto tr182
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st160
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr182
		}
		goto st0
tr182:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 161; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st161
	st161:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof161
		}
	st_case_161:
//line ./combat/tokenizer.go:3476
		switch ( t.data)[( t.p)] {
		case 32:
			goto st161
		case 45:
			goto tr184
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr185
			}
		case ( t.data)[( t.p)] >= 9:
			goto st161
		}
		goto st0
tr184:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st162
	st162:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof162
		}
	st_case_162:
//line ./combat/tokenizer.go:3506
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st163
		}
		goto st0
tr185:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:3525
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr187
		case 46:
			goto st205
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st163
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr187
		}
		goto st0
tr187:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 164; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st164
	st164:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof164
		}
	st_case_164:
//line ./combat/tokenizer.go:3557
		switch ( t.data)[( t.p)] {
		case 32:
			goto st164
		case 40:
			goto st165
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st164
		}
		goto st0
	st165:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof165
		}
	st_case_165:
		if ( t.data)[( t.p)] == 104 {
			goto st166
		}
		goto st0
	st166:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof166
		}
	st_case_166:
		if ( t.data)[( t.p)] == 58 {
			goto st167
		}
		goto st0
	st167:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof167
		}
	st_case_167:
		if ( t.data)[( t.p)] == 45 {
			goto tr193
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr194
		}
		goto st0
tr193:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st168
	st168:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof168
		}
	st_case_168:
//line ./combat/tokenizer.go:3612
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st169
		}
		goto st0
tr194:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st169
	st169:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof169
		}
	st_case_169:
//line ./combat/tokenizer.go:3631
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr196
		case 46:
			goto st204
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st169
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr196
		}
		goto st0
tr196:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 170; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st170
	st170:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof170
		}
	st_case_170:
//line ./combat/tokenizer.go:3663
		switch ( t.data)[( t.p)] {
		case 32:
			goto st170
		case 115:
			goto st171
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st170
		}
		goto st0
	st171:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof171
		}
	st_case_171:
		if ( t.data)[( t.p)] == 58 {
			goto st172
		}
		goto st0
	st172:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof172
		}
	st_case_172:
		if ( t.data)[( t.p)] == 45 {
			goto tr201
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr202
		}
		goto st0
tr201:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st173
	st173:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof173
		}
	st_case_173:
//line ./combat/tokenizer.go:3709
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st174
		}
		goto st0
tr202:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:3728
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr204
		case 46:
			goto st203
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st174
		}
		goto st0
tr204:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 175; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st175
	st175:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof175
		}
	st_case_175:
//line ./combat/tokenizer.go:3755
		if ( t.data)[( t.p)] == 32 {
			goto st176
		}
		goto st0
	st176:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st177
		case 40:
			goto tr208
		case 95:
			goto tr209
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr209
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr209
			}
		default:
			goto tr209
		}
		goto st0
tr235:
//line ./combat/tokenizer.go.rl:137
t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))
	goto st177
tr594:
//line ./combat/tokenizer.go.rl:138
t.tokval(newAnyVal(DAMAGE_MODIFIER, strTok(t.data[t.prev:t.p])))
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st177
	st177:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof177
		}
	st_case_177:
//line ./combat/tokenizer.go:3801
		if ( t.data)[( t.p)] == 95 {
			goto tr210
		}
		if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
			goto tr210
		}
		goto st0
tr210:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st523
	st523:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof523
		}
	st_case_523:
//line ./combat/tokenizer.go:3823
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr592
		case 95:
			goto st523
		case 124:
			goto tr594
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 65 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 90 {
				goto st523
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr592
		}
		goto st0
tr592:
//line ./combat/tokenizer.go.rl:138
t.tokval(newAnyVal(DAMAGE_MODIFIER, strTok(t.data[t.prev:t.p])))
	goto st524
	st524:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof524
		}
	st_case_524:
//line ./combat/tokenizer.go:3850
		switch ( t.data)[( t.p)] {
		case 32:
			goto st524
		case 60:
			goto st178
		case 82:
			goto tr597
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st524
		}
		goto st0
	st178:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof178
		}
	st_case_178:
		if ( t.data)[( t.p)] == 70 {
			goto st179
		}
		goto st0
	st179:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof179
		}
	st_case_179:
		if ( t.data)[( t.p)] == 114 {
			goto st180
		}
		goto st0
	st180:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof180
		}
	st_case_180:
		if ( t.data)[( t.p)] == 105 {
			goto st181
		}
		goto st0
	st181:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof181
		}
	st_case_181:
		if ( t.data)[( t.p)] == 101 {
			goto st182
		}
		goto st0
	st182:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof182
		}
	st_case_182:
		if ( t.data)[( t.p)] == 110 {
			goto st183
		}
		goto st0
	st183:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof183
		}
	st_case_183:
		if ( t.data)[( t.p)] == 100 {
			goto st184
		}
		goto st0
	st184:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof184
		}
	st_case_184:
		if ( t.data)[( t.p)] == 108 {
			goto st185
		}
		goto st0
	st185:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof185
		}
	st_case_185:
		if ( t.data)[( t.p)] == 121 {
			goto st186
		}
		goto st0
	st186:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof186
		}
	st_case_186:
		if ( t.data)[( t.p)] == 70 {
			goto st187
		}
		goto st0
	st187:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof187
		}
	st_case_187:
		if ( t.data)[( t.p)] == 105 {
			goto st188
		}
		goto st0
	st188:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof188
		}
	st_case_188:
		if ( t.data)[( t.p)] == 114 {
			goto st189
		}
		goto st0
	st189:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof189
		}
	st_case_189:
		if ( t.data)[( t.p)] == 101 {
			goto st190
		}
		goto st0
	st190:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof190
		}
	st_case_190:
		if ( t.data)[( t.p)] == 62 {
			goto tr223
		}
		goto st0
tr223:
//line ./combat/tokenizer.go.rl:146
t.tok(FRIENDLY_FIRE)
	goto st525
	st525:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof525
		}
	st_case_525:
//line ./combat/tokenizer.go:3989
		if ( t.data)[( t.p)] == 32 {
			goto st526
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st526
		}
		goto st0
	st526:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof526
		}
	st_case_526:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st526
		case 82:
			goto tr597
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st526
		}
		goto st0
tr597:
//line ./combat/tokenizer.go.rl:147
t.tok(ROCKET)
	goto st191
	st191:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof191
		}
	st_case_191:
//line ./combat/tokenizer.go:4021
		if ( t.data)[( t.p)] == 111 {
			goto st192
		}
		goto st0
	st192:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof192
		}
	st_case_192:
		if ( t.data)[( t.p)] == 99 {
			goto st193
		}
		goto st0
	st193:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof193
		}
	st_case_193:
		if ( t.data)[( t.p)] == 107 {
			goto st194
		}
		goto st0
	st194:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof194
		}
	st_case_194:
		if ( t.data)[( t.p)] == 101 {
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
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st197
		}
		goto st0
	st197:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof197
		}
	st_case_197:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st197
		case 45:
			goto tr230
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr231
			}
		case ( t.data)[( t.p)] >= 9:
			goto st197
		}
		goto st0
tr230:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st198
	st198:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof198
		}
	st_case_198:
//line ./combat/tokenizer.go:4108
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st527
		}
		goto st0
tr231:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st527
	st527:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof527
		}
	st_case_527:
//line ./combat/tokenizer.go:4127
		if ( t.data)[( t.p)] == 32 {
			goto tr599
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st527
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr599
		}
		goto st0
tr208:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st199
	st199:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof199
		}
	st_case_199:
//line ./combat/tokenizer.go:4154
		if ( t.data)[( t.p)] == 95 {
			goto st200
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st200
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st200
			}
		default:
			goto st200
		}
		goto st0
	st200:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof200
		}
	st_case_200:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st201
		case 95:
			goto st200
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st200
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st200
			}
		default:
			goto st200
		}
		goto st0
	st201:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof201
		}
	st_case_201:
		if ( t.data)[( t.p)] == 32 {
			goto tr235
		}
		goto st0
tr209:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st202
	st202:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof202
		}
	st_case_202:
//line ./combat/tokenizer.go:4218
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr235
		case 95:
			goto st202
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st202
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st202
			}
		default:
			goto st202
		}
		goto st0
	st203:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof203
		}
	st_case_203:
		if ( t.data)[( t.p)] == 41 {
			goto tr204
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st203
		}
		goto st0
	st204:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof204
		}
	st_case_204:
		if ( t.data)[( t.p)] == 32 {
			goto tr196
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st204
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr196
		}
		goto st0
	st205:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof205
		}
	st_case_205:
		if ( t.data)[( t.p)] == 32 {
			goto tr187
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st205
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr187
		}
		goto st0
tr174:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st206
	st206:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof206
		}
	st_case_206:
//line ./combat/tokenizer.go:4298
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr177
		case 95:
			goto st206
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st206
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st206
			}
		default:
			goto st206
		}
		goto st0
tr168:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st207
	st207:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof207
		}
	st_case_207:
//line ./combat/tokenizer.go:4332
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr171
		case 95:
			goto st207
		case 124:
			goto tr172
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st207
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st207
			}
		default:
			goto st207
		}
		goto st0
tr154:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st208
	st208:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof208
		}
	st_case_208:
//line ./combat/tokenizer.go:4368
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr157
		case 95:
			goto st208
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st208
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st208
			}
		default:
			goto st208
		}
		goto st0
tr148:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st209
	st209:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof209
		}
	st_case_209:
//line ./combat/tokenizer.go:4402
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr151
		case 95:
			goto st209
		case 124:
			goto tr152
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st209
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st209
			}
		default:
			goto st209
		}
		goto st0
tr24:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st210
	st210:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof210
		}
	st_case_210:
//line ./combat/tokenizer.go:4438
		if ( t.data)[( t.p)] == 97 {
			goto st211
		}
		goto st0
	st211:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof211
		}
	st_case_211:
		if ( t.data)[( t.p)] == 109 {
			goto st212
		}
		goto st0
	st212:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof212
		}
	st_case_212:
		if ( t.data)[( t.p)] == 101 {
			goto st213
		}
		goto st0
	st213:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof213
		}
	st_case_213:
		if ( t.data)[( t.p)] == 112 {
			goto st214
		}
		goto st0
	st214:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof214
		}
	st_case_214:
		if ( t.data)[( t.p)] == 108 {
			goto st215
		}
		goto st0
	st215:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof215
		}
	st_case_215:
		if ( t.data)[( t.p)] == 97 {
			goto st216
		}
		goto st0
	st216:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof216
		}
	st_case_216:
		if ( t.data)[( t.p)] == 121 {
			goto st217
		}
		goto st0
	st217:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof217
		}
	st_case_217:
		if ( t.data)[( t.p)] == 32 {
			goto st218
		}
		goto st0
	st218:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof218
		}
	st_case_218:
		if ( t.data)[( t.p)] == 102 {
			goto st219
		}
		goto st0
	st219:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof219
		}
	st_case_219:
		if ( t.data)[( t.p)] == 105 {
			goto st220
		}
		goto st0
	st220:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof220
		}
	st_case_220:
		if ( t.data)[( t.p)] == 110 {
			goto st221
		}
		goto st0
	st221:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof221
		}
	st_case_221:
		if ( t.data)[( t.p)] == 105 {
			goto st222
		}
		goto st0
	st222:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof222
		}
	st_case_222:
		if ( t.data)[( t.p)] == 115 {
			goto st223
		}
		goto st0
	st223:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof223
		}
	st_case_223:
		if ( t.data)[( t.p)] == 104 {
			goto st224
		}
		goto st0
	st224:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof224
		}
	st_case_224:
		if ( t.data)[( t.p)] == 101 {
			goto st225
		}
		goto st0
	st225:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof225
		}
	st_case_225:
		if ( t.data)[( t.p)] == 100 {
			goto st226
		}
		goto st0
	st226:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof226
		}
	st_case_226:
		if ( t.data)[( t.p)] == 46 {
			goto tr257
		}
		goto st0
tr257:
//line ./combat/tokenizer.go.rl:130
 t.tok(GAMEPLAY_FINISHED) 
	goto st227
	st227:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof227
		}
	st_case_227:
//line ./combat/tokenizer.go:4596
		if ( t.data)[( t.p)] == 32 {
			goto st228
		}
		goto st0
	st228:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof228
		}
	st_case_228:
		if ( t.data)[( t.p)] == 87 {
			goto st229
		}
		goto st0
	st229:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof229
		}
	st_case_229:
		if ( t.data)[( t.p)] == 105 {
			goto st230
		}
		goto st0
	st230:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof230
		}
	st_case_230:
		if ( t.data)[( t.p)] == 110 {
			goto st231
		}
		goto st0
	st231:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof231
		}
	st_case_231:
		if ( t.data)[( t.p)] == 110 {
			goto st232
		}
		goto st0
	st232:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof232
		}
	st_case_232:
		if ( t.data)[( t.p)] == 101 {
			goto st233
		}
		goto st0
	st233:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof233
		}
	st_case_233:
		if ( t.data)[( t.p)] == 114 {
			goto st234
		}
		goto st0
	st234:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof234
		}
	st_case_234:
		if ( t.data)[( t.p)] == 32 {
			goto st235
		}
		goto st0
	st235:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof235
		}
	st_case_235:
		if ( t.data)[( t.p)] == 116 {
			goto st236
		}
		goto st0
	st236:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof236
		}
	st_case_236:
		if ( t.data)[( t.p)] == 101 {
			goto st237
		}
		goto st0
	st237:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof237
		}
	st_case_237:
		if ( t.data)[( t.p)] == 97 {
			goto st238
		}
		goto st0
	st238:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof238
		}
	st_case_238:
		if ( t.data)[( t.p)] == 109 {
			goto st239
		}
		goto st0
	st239:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof239
		}
	st_case_239:
		if ( t.data)[( t.p)] == 58 {
			goto st240
		}
		goto st0
	st240:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof240
		}
	st_case_240:
		if ( t.data)[( t.p)] == 32 {
			goto st241
		}
		goto st0
	st241:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof241
		}
	st_case_241:
		if ( t.data)[( t.p)] == 45 {
			goto tr272
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr273
		}
		goto st0
tr272:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:4744
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st243
		}
		goto st0
tr273:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st243
	st243:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof243
		}
	st_case_243:
//line ./combat/tokenizer.go:4763
		if ( t.data)[( t.p)] == 40 {
			goto tr275
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st243
		}
		goto st0
tr275:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 244; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st244
	st244:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof244
		}
	st_case_244:
//line ./combat/tokenizer.go:4787
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr276
		case 95:
			goto tr277
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr277
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr277
			}
		default:
			goto tr277
		}
		goto st0
tr276:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st245
	st245:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof245
		}
	st_case_245:
//line ./combat/tokenizer.go:4821
		if ( t.data)[( t.p)] == 95 {
			goto st246
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st246
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st246
			}
		default:
			goto st246
		}
		goto st0
	st246:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st247
		case 95:
			goto st246
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st246
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st246
			}
		default:
			goto st246
		}
		goto st0
	st247:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof247
		}
	st_case_247:
		if ( t.data)[( t.p)] == 41 {
			goto tr280
		}
		goto st0
tr280:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st248
	st248:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof248
		}
	st_case_248:
//line ./combat/tokenizer.go:4885
		if ( t.data)[( t.p)] == 46 {
			goto st249
		}
		goto st0
	st249:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof249
		}
	st_case_249:
		if ( t.data)[( t.p)] == 32 {
			goto st250
		}
		goto st0
	st250:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof250
		}
	st_case_250:
		if ( t.data)[( t.p)] == 70 {
			goto st251
		}
		goto st0
	st251:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof251
		}
	st_case_251:
		if ( t.data)[( t.p)] == 105 {
			goto st252
		}
		goto st0
	st252:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof252
		}
	st_case_252:
		if ( t.data)[( t.p)] == 110 {
			goto st253
		}
		goto st0
	st253:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof253
		}
	st_case_253:
		if ( t.data)[( t.p)] == 105 {
			goto st254
		}
		goto st0
	st254:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof254
		}
	st_case_254:
		if ( t.data)[( t.p)] == 115 {
			goto st255
		}
		goto st0
	st255:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof255
		}
	st_case_255:
		if ( t.data)[( t.p)] == 104 {
			goto st256
		}
		goto st0
	st256:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof256
		}
	st_case_256:
		if ( t.data)[( t.p)] == 32 {
			goto st257
		}
		goto st0
	st257:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof257
		}
	st_case_257:
		if ( t.data)[( t.p)] == 114 {
			goto st258
		}
		goto st0
	st258:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof258
		}
	st_case_258:
		if ( t.data)[( t.p)] == 101 {
			goto st259
		}
		goto st0
	st259:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof259
		}
	st_case_259:
		if ( t.data)[( t.p)] == 97 {
			goto st260
		}
		goto st0
	st260:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof260
		}
	st_case_260:
		if ( t.data)[( t.p)] == 115 {
			goto st261
		}
		goto st0
	st261:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof261
		}
	st_case_261:
		if ( t.data)[( t.p)] == 111 {
			goto st262
		}
		goto st0
	st262:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof262
		}
	st_case_262:
		if ( t.data)[( t.p)] == 110 {
			goto st263
		}
		goto st0
	st263:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof263
		}
	st_case_263:
		if ( t.data)[( t.p)] == 58 {
			goto st264
		}
		goto st0
	st264:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof264
		}
	st_case_264:
		if ( t.data)[( t.p)] == 32 {
			goto st265
		}
		goto st0
	st265:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof265
		}
	st_case_265:
		if ( t.data)[( t.p)] == 39 {
			goto st266
		}
		goto st0
	st266:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof266
		}
	st_case_266:
		if ( t.data)[( t.p)] == 95 {
			goto tr299
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr299
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr299
			}
		default:
			goto tr299
		}
		goto st0
tr299:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st267
	st267:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof267
		}
	st_case_267:
//line ./combat/tokenizer.go:5078
		switch ( t.data)[( t.p)] {
		case 32:
			goto st268
		case 39:
			goto tr301
		case 95:
			goto st267
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st267
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st267
			}
		default:
			goto st267
		}
		goto st0
	st268:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof268
		}
	st_case_268:
		if ( t.data)[( t.p)] == 95 {
			goto st267
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st267
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st267
			}
		default:
			goto st267
		}
		goto st0
tr301:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st269
	st269:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof269
		}
	st_case_269:
//line ./combat/tokenizer.go:5135
		if ( t.data)[( t.p)] == 46 {
			goto st270
		}
		goto st0
	st270:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof270
		}
	st_case_270:
		if ( t.data)[( t.p)] == 32 {
			goto st271
		}
		goto st0
	st271:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof271
		}
	st_case_271:
		if ( t.data)[( t.p)] == 65 {
			goto st272
		}
		goto st0
	st272:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof272
		}
	st_case_272:
		if ( t.data)[( t.p)] == 99 {
			goto st273
		}
		goto st0
	st273:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof273
		}
	st_case_273:
		if ( t.data)[( t.p)] == 116 {
			goto st274
		}
		goto st0
	st274:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof274
		}
	st_case_274:
		if ( t.data)[( t.p)] == 117 {
			goto st275
		}
		goto st0
	st275:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof275
		}
	st_case_275:
		if ( t.data)[( t.p)] == 97 {
			goto st276
		}
		goto st0
	st276:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof276
		}
	st_case_276:
		if ( t.data)[( t.p)] == 108 {
			goto st277
		}
		goto st0
	st277:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof277
		}
	st_case_277:
		if ( t.data)[( t.p)] == 32 {
			goto st278
		}
		goto st0
	st278:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof278
		}
	st_case_278:
		if ( t.data)[( t.p)] == 103 {
			goto st279
		}
		goto st0
	st279:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof279
		}
	st_case_279:
		if ( t.data)[( t.p)] == 97 {
			goto st280
		}
		goto st0
	st280:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof280
		}
	st_case_280:
		if ( t.data)[( t.p)] == 109 {
			goto st281
		}
		goto st0
	st281:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof281
		}
	st_case_281:
		if ( t.data)[( t.p)] == 101 {
			goto st282
		}
		goto st0
	st282:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof282
		}
	st_case_282:
		if ( t.data)[( t.p)] == 32 {
			goto st283
		}
		goto st0
	st283:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof283
		}
	st_case_283:
		if ( t.data)[( t.p)] == 116 {
			goto st284
		}
		goto st0
	st284:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof284
		}
	st_case_284:
		if ( t.data)[( t.p)] == 105 {
			goto st285
		}
		goto st0
	st285:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof285
		}
	st_case_285:
		if ( t.data)[( t.p)] == 109 {
			goto st286
		}
		goto st0
	st286:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof286
		}
	st_case_286:
		if ( t.data)[( t.p)] == 101 {
			goto st287
		}
		goto st0
	st287:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof287
		}
	st_case_287:
		if ( t.data)[( t.p)] == 32 {
			goto st288
		}
		goto st0
	st288:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof288
		}
	st_case_288:
		if ( t.data)[( t.p)] == 45 {
			goto tr322
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr323
		}
		goto st0
tr322:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st289
	st289:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof289
		}
	st_case_289:
//line ./combat/tokenizer.go:5328
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st290
		}
		goto st0
tr323:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st290
	st290:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof290
		}
	st_case_290:
//line ./combat/tokenizer.go:5347
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr325
		case 46:
			goto st294
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st290
		}
		goto st0
tr325:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 291; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st291
	st291:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof291
		}
	st_case_291:
//line ./combat/tokenizer.go:5374
		if ( t.data)[( t.p)] == 115 {
			goto st292
		}
		goto st0
	st292:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof292
		}
	st_case_292:
		if ( t.data)[( t.p)] == 101 {
			goto st293
		}
		goto st0
	st293:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof293
		}
	st_case_293:
		if ( t.data)[( t.p)] == 99 {
			goto st522
		}
		goto st0
	st294:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof294
		}
	st_case_294:
		if ( t.data)[( t.p)] == 32 {
			goto tr325
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st294
		}
		goto st0
tr277:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st295
	st295:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof295
		}
	st_case_295:
//line ./combat/tokenizer.go:5423
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr280
		case 95:
			goto st295
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st295
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st295
			}
		default:
			goto st295
		}
		goto st0
tr25:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:5457
		if ( t.data)[( t.p)] == 101 {
			goto st297
		}
		goto st0
	st297:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof297
		}
	st_case_297:
		if ( t.data)[( t.p)] == 97 {
			goto st298
		}
		goto st0
	st298:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof298
		}
	st_case_298:
		if ( t.data)[( t.p)] == 108 {
			goto tr332
		}
		goto st0
tr332:
//line ./combat/tokenizer.go.rl:126
 t.tok(HEAL) 
	goto st299
	st299:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof299
		}
	st_case_299:
//line ./combat/tokenizer.go:5489
		if ( t.data)[( t.p)] == 32 {
			goto st300
		}
		goto st0
	st300:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof300
		}
	st_case_300:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st300
		case 40:
			goto tr334
		case 95:
			goto tr335
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr335
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr335
			}
		default:
			goto tr335
		}
		goto st0
tr334:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st301
	st301:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof301
		}
	st_case_301:
//line ./combat/tokenizer.go:5534
		if ( t.data)[( t.p)] == 95 {
			goto st302
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st302
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st302
			}
		default:
			goto st302
		}
		goto st0
	st302:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof302
		}
	st_case_302:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st303
		case 95:
			goto st302
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st302
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st302
			}
		default:
			goto st302
		}
		goto st0
	st303:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof303
		}
	st_case_303:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr338
		case 124:
			goto tr339
		}
		goto st0
tr338:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st304
	st304:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof304
		}
	st_case_304:
//line ./combat/tokenizer.go:5603
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr340
		case 95:
			goto tr341
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr341
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr341
			}
		default:
			goto tr341
		}
		goto st0
tr340:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st305
	st305:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof305
		}
	st_case_305:
//line ./combat/tokenizer.go:5637
		if ( t.data)[( t.p)] == 95 {
			goto st306
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st306
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st306
			}
		default:
			goto st306
		}
		goto st0
	st306:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st307
		case 95:
			goto st306
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st306
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st306
			}
		default:
			goto st306
		}
		goto st0
	st307:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof307
		}
	st_case_307:
		if ( t.data)[( t.p)] == 41 {
			goto tr344
		}
		goto st0
tr344:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st308
	st308:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof308
		}
	st_case_308:
//line ./combat/tokenizer.go:5703
		if ( t.data)[( t.p)] == 124 {
			goto tr345
		}
		goto st0
tr339:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st309
tr345:
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st309
	st309:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof309
		}
	st_case_309:
//line ./combat/tokenizer.go:5728
		if ( t.data)[( t.p)] == 45 {
			goto tr346
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr347
		}
		goto st0
tr346:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st310
	st310:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof310
		}
	st_case_310:
//line ./combat/tokenizer.go:5750
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st311
		}
		goto st0
tr347:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:5769
		if ( t.data)[( t.p)] == 32 {
			goto tr349
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st311
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr349
		}
		goto st0
tr349:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 312; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st312
	st312:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof312
		}
	st_case_312:
//line ./combat/tokenizer.go:5798
		switch ( t.data)[( t.p)] {
		case 32:
			goto st312
		case 45:
			goto st313
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st312
		}
		goto st0
	st313:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof313
		}
	st_case_313:
		if ( t.data)[( t.p)] == 62 {
			goto tr352
		}
		goto st0
tr352:
//line ./combat/tokenizer.go.rl:143
t.tok(ARROW)
	goto st314
	st314:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof314
		}
	st_case_314:
//line ./combat/tokenizer.go:5827
		if ( t.data)[( t.p)] == 32 {
			goto st315
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st315
		}
		goto st0
	st315:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st315
		case 40:
			goto tr354
		case 95:
			goto tr355
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st315
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr355
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr355
			}
		default:
			goto tr355
		}
		goto st0
tr354:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st316
	st316:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof316
		}
	st_case_316:
//line ./combat/tokenizer.go:5880
		if ( t.data)[( t.p)] == 95 {
			goto st317
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st317
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st317
			}
		default:
			goto st317
		}
		goto st0
	st317:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st318
		case 95:
			goto st317
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st317
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st317
			}
		default:
			goto st317
		}
		goto st0
	st318:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof318
		}
	st_case_318:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr358
		case 124:
			goto tr359
		}
		goto st0
tr358:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st319
	st319:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof319
		}
	st_case_319:
//line ./combat/tokenizer.go:5949
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr360
		case 95:
			goto tr361
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr361
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr361
			}
		default:
			goto tr361
		}
		goto st0
tr360:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st320
	st320:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof320
		}
	st_case_320:
//line ./combat/tokenizer.go:5983
		if ( t.data)[( t.p)] == 95 {
			goto st321
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st321
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st321
			}
		default:
			goto st321
		}
		goto st0
	st321:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st322
		case 95:
			goto st321
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st321
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st321
			}
		default:
			goto st321
		}
		goto st0
	st322:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof322
		}
	st_case_322:
		if ( t.data)[( t.p)] == 41 {
			goto tr364
		}
		goto st0
tr364:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st323
	st323:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof323
		}
	st_case_323:
//line ./combat/tokenizer.go:6049
		if ( t.data)[( t.p)] == 124 {
			goto tr365
		}
		goto st0
tr359:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st324
tr365:
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st324
	st324:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof324
		}
	st_case_324:
//line ./combat/tokenizer.go:6074
		if ( t.data)[( t.p)] == 45 {
			goto tr366
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr367
		}
		goto st0
tr366:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st325
	st325:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof325
		}
	st_case_325:
//line ./combat/tokenizer.go:6096
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st326
		}
		goto st0
tr367:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st326
	st326:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof326
		}
	st_case_326:
//line ./combat/tokenizer.go:6115
		if ( t.data)[( t.p)] == 32 {
			goto tr369
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st326
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr369
		}
		goto st0
tr369:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 327; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st327
	st327:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof327
		}
	st_case_327:
//line ./combat/tokenizer.go:6144
		switch ( t.data)[( t.p)] {
		case 32:
			goto st327
		case 45:
			goto tr371
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr372
			}
		case ( t.data)[( t.p)] >= 9:
			goto st327
		}
		goto st0
tr371:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st328
	st328:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof328
		}
	st_case_328:
//line ./combat/tokenizer.go:6174
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st329
		}
		goto st0
tr372:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st329
	st329:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof329
		}
	st_case_329:
//line ./combat/tokenizer.go:6193
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr374
		case 46:
			goto st333
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st329
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr374
		}
		goto st0
tr374:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 330; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st330
	st330:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof330
		}
	st_case_330:
//line ./combat/tokenizer.go:6225
		switch ( t.data)[( t.p)] {
		case 32:
			goto st330
		case 40:
			goto tr377
		case 95:
			goto tr378
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st330
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr378
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr378
			}
		default:
			goto tr378
		}
		goto st0
tr377:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st331
	st331:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof331
		}
	st_case_331:
//line ./combat/tokenizer.go:6266
		if ( t.data)[( t.p)] == 95 {
			goto st332
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st332
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st332
			}
		default:
			goto st332
		}
		goto st0
	st332:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st528
		case 95:
			goto st332
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st332
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st332
			}
		default:
			goto st332
		}
		goto st0
	st528:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof528
		}
	st_case_528:
		if ( t.data)[( t.p)] == 32 {
			goto tr600
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr600
		}
		goto st0
tr378:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:6333
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr600
		case 95:
			goto st529
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr600
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st529
				}
			case ( t.data)[( t.p)] >= 65:
				goto st529
			}
		default:
			goto st529
		}
		goto st0
	st333:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof333
		}
	st_case_333:
		if ( t.data)[( t.p)] == 32 {
			goto tr374
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st333
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr374
		}
		goto st0
tr361:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st334
	st334:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof334
		}
	st_case_334:
//line ./combat/tokenizer.go:6389
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr364
		case 95:
			goto st334
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st334
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st334
			}
		default:
			goto st334
		}
		goto st0
tr355:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st335
	st335:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof335
		}
	st_case_335:
//line ./combat/tokenizer.go:6423
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr358
		case 95:
			goto st335
		case 124:
			goto tr359
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st335
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st335
			}
		default:
			goto st335
		}
		goto st0
tr341:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:6459
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr344
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
tr335:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st337
	st337:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof337
		}
	st_case_337:
//line ./combat/tokenizer.go:6493
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr338
		case 95:
			goto st337
		case 124:
			goto tr339
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st337
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st337
			}
		default:
			goto st337
		}
		goto st0
tr26:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st338
	st338:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof338
		}
	st_case_338:
//line ./combat/tokenizer.go:6529
		if ( t.data)[( t.p)] == 105 {
			goto st339
		}
		goto st0
	st339:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof339
		}
	st_case_339:
		if ( t.data)[( t.p)] == 108 {
			goto st340
		}
		goto st0
	st340:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof340
		}
	st_case_340:
		if ( t.data)[( t.p)] == 108 {
			goto st341
		}
		goto st0
	st341:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof341
		}
	st_case_341:
		if ( t.data)[( t.p)] == 101 {
			goto st342
		}
		goto st0
	st342:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof342
		}
	st_case_342:
		if ( t.data)[( t.p)] == 100 {
			goto tr389
		}
		goto st0
tr389:
//line ./combat/tokenizer.go.rl:127
 t.tok(KILL) 
	goto st343
	st343:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof343
		}
	st_case_343:
//line ./combat/tokenizer.go:6579
		if ( t.data)[( t.p)] == 32 {
			goto st344
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st344
		}
		goto st0
	st344:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st344
		case 40:
			goto tr391
		case 95:
			goto tr392
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st344
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr392
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr392
			}
		default:
			goto tr392
		}
		goto st0
tr391:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st345
	st345:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof345
		}
	st_case_345:
//line ./combat/tokenizer.go:6632
		if ( t.data)[( t.p)] == 95 {
			goto st346
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st346
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st346
			}
		default:
			goto st346
		}
		goto st0
	st346:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof346
		}
	st_case_346:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st347
		case 95:
			goto st346
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st346
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st346
			}
		default:
			goto st346
		}
		goto st0
	st347:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof347
		}
	st_case_347:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr395
		case 40:
			goto tr396
		case 124:
			goto tr397
		}
		goto st0
tr395:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st348
	st348:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof348
		}
	st_case_348:
//line ./combat/tokenizer.go:6703
		switch ( t.data)[( t.p)] {
		case 32:
			goto st348
		case 40:
			goto tr399
		case 95:
			goto tr400
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st348
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr400
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr400
			}
		default:
			goto tr400
		}
		goto st0
tr399:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st349
	st349:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof349
		}
	st_case_349:
//line ./combat/tokenizer.go:6744
		if ( t.data)[( t.p)] == 95 {
			goto st350
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st350
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st350
			}
		default:
			goto st350
		}
		goto st0
	st350:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st351
		case 95:
			goto st350
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st350
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st350
			}
		default:
			goto st350
		}
		goto st0
	st351:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr396
		case 124:
			goto tr397
		}
		goto st0
tr396:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st352
	st352:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof352
		}
	st_case_352:
//line ./combat/tokenizer.go:6813
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr403
		case 95:
			goto tr404
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr404
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr404
			}
		default:
			goto tr404
		}
		goto st0
tr403:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st353
	st353:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof353
		}
	st_case_353:
//line ./combat/tokenizer.go:6847
		if ( t.data)[( t.p)] == 95 {
			goto st354
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st354
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st354
			}
		default:
			goto st354
		}
		goto st0
	st354:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof354
		}
	st_case_354:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st355
		case 95:
			goto st354
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st354
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st354
			}
		default:
			goto st354
		}
		goto st0
	st355:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof355
		}
	st_case_355:
		if ( t.data)[( t.p)] == 41 {
			goto tr407
		}
		goto st0
tr407:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st356
	st356:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof356
		}
	st_case_356:
//line ./combat/tokenizer.go:6913
		if ( t.data)[( t.p)] == 124 {
			goto tr408
		}
		goto st0
tr397:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st357
tr408:
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st357
	st357:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof357
		}
	st_case_357:
//line ./combat/tokenizer.go:6938
		if ( t.data)[( t.p)] == 45 {
			goto tr409
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr410
		}
		goto st0
tr409:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st358
	st358:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof358
		}
	st_case_358:
//line ./combat/tokenizer.go:6960
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st359
		}
		goto st0
tr410:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st359
	st359:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof359
		}
	st_case_359:
//line ./combat/tokenizer.go:6979
		if ( t.data)[( t.p)] == 59 {
			goto tr412
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st359
		}
		goto st0
tr412:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 360; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st360
	st360:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof360
		}
	st_case_360:
//line ./combat/tokenizer.go:7005
		if ( t.data)[( t.p)] == 32 {
			goto st361
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st361
		}
		goto st0
	st361:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof361
		}
	st_case_361:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st361
		case 107:
			goto st362
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st361
		}
		goto st0
	st362:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof362
		}
	st_case_362:
		if ( t.data)[( t.p)] == 105 {
			goto st363
		}
		goto st0
	st363:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof363
		}
	st_case_363:
		if ( t.data)[( t.p)] == 108 {
			goto st364
		}
		goto st0
	st364:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof364
		}
	st_case_364:
		if ( t.data)[( t.p)] == 108 {
			goto st365
		}
		goto st0
	st365:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof365
		}
	st_case_365:
		if ( t.data)[( t.p)] == 101 {
			goto st366
		}
		goto st0
	st366:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof366
		}
	st_case_366:
		if ( t.data)[( t.p)] == 114 {
			goto st367
		}
		goto st0
	st367:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof367
		}
	st_case_367:
		if ( t.data)[( t.p)] == 32 {
			goto st368
		}
		goto st0
	st368:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr421
		case 95:
			goto tr422
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr422
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr422
			}
		default:
			goto tr422
		}
		goto st0
tr421:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:7120
		if ( t.data)[( t.p)] == 95 {
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
	st370:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st371
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
	st371:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof371
		}
	st_case_371:
		if ( t.data)[( t.p)] == 124 {
			goto tr425
		}
		goto st0
tr425:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line ./combat/tokenizer.go.rl:117
 t.tok(int(t.data[t.p]))
	goto st372
	st372:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof372
		}
	st_case_372:
//line ./combat/tokenizer.go:7186
		if ( t.data)[( t.p)] == 45 {
			goto tr426
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr427
		}
		goto st0
tr426:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st373
	st373:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof373
		}
	st_case_373:
//line ./combat/tokenizer.go:7208
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st374
		}
		goto st0
tr427:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st374
	st374:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof374
		}
	st_case_374:
//line ./combat/tokenizer.go:7227
		if ( t.data)[( t.p)] == 32 {
			goto tr429
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st374
		}
		goto st0
tr429:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 530; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st530
	st530:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof530
		}
	st_case_530:
//line ./combat/tokenizer.go:7251
		switch ( t.data)[( t.p)] {
		case 32:
			goto st531
		case 40:
			goto tr603
		case 95:
			goto tr604
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st522
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr604
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr604
			}
		default:
			goto tr604
		}
		goto st0
tr606:
//line ./combat/tokenizer.go.rl:137
t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))
	goto st531
	st531:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof531
		}
	st_case_531:
//line ./combat/tokenizer.go:7287
		switch ( t.data)[( t.p)] {
		case 32:
			goto st522
		case 60:
			goto st375
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
tr613:
//line ./combat/tokenizer.go.rl:149
t.tok(PARTICIPATION_MODIFIERS_END)
	goto st375
	st375:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof375
		}
	st_case_375:
//line ./combat/tokenizer.go:7307
		if ( t.data)[( t.p)] == 70 {
			goto st376
		}
		goto st0
	st376:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof376
		}
	st_case_376:
		if ( t.data)[( t.p)] == 114 {
			goto st377
		}
		goto st0
	st377:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof377
		}
	st_case_377:
		if ( t.data)[( t.p)] == 105 {
			goto st378
		}
		goto st0
	st378:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof378
		}
	st_case_378:
		if ( t.data)[( t.p)] == 101 {
			goto st379
		}
		goto st0
	st379:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof379
		}
	st_case_379:
		if ( t.data)[( t.p)] == 110 {
			goto st380
		}
		goto st0
	st380:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof380
		}
	st_case_380:
		if ( t.data)[( t.p)] == 100 {
			goto st381
		}
		goto st0
	st381:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof381
		}
	st_case_381:
		if ( t.data)[( t.p)] == 108 {
			goto st382
		}
		goto st0
	st382:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof382
		}
	st_case_382:
		if ( t.data)[( t.p)] == 121 {
			goto st383
		}
		goto st0
	st383:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof383
		}
	st_case_383:
		if ( t.data)[( t.p)] == 70 {
			goto st384
		}
		goto st0
	st384:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof384
		}
	st_case_384:
		if ( t.data)[( t.p)] == 105 {
			goto st385
		}
		goto st0
	st385:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof385
		}
	st_case_385:
		if ( t.data)[( t.p)] == 114 {
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
		if ( t.data)[( t.p)] == 62 {
			goto tr442
		}
		goto st0
tr603:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st388
	st388:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof388
		}
	st_case_388:
//line ./combat/tokenizer.go:7434
		if ( t.data)[( t.p)] == 95 {
			goto st389
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st389
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st389
			}
		default:
			goto st389
		}
		goto st0
	st389:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof389
		}
	st_case_389:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st532
		case 95:
			goto st389
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st389
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st389
			}
		default:
			goto st389
		}
		goto st0
	st532:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof532
		}
	st_case_532:
		if ( t.data)[( t.p)] == 32 {
			goto tr606
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr600
		}
		goto st0
tr604:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st533
	st533:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof533
		}
	st_case_533:
//line ./combat/tokenizer.go:7501
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr606
		case 95:
			goto st533
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr600
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st533
				}
			case ( t.data)[( t.p)] >= 65:
				goto st533
			}
		default:
			goto st533
		}
		goto st0
tr422:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st390
	st390:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof390
		}
	st_case_390:
//line ./combat/tokenizer.go:7540
		switch ( t.data)[( t.p)] {
		case 95:
			goto st390
		case 124:
			goto tr425
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st390
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st390
			}
		default:
			goto st390
		}
		goto st0
tr404:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st391
	st391:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof391
		}
	st_case_391:
//line ./combat/tokenizer.go:7574
		switch ( t.data)[( t.p)] {
		case 41:
			goto tr407
		case 95:
			goto st391
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st391
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st391
			}
		default:
			goto st391
		}
		goto st0
tr400:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st392
	st392:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof392
		}
	st_case_392:
//line ./combat/tokenizer.go:7608
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr396
		case 95:
			goto st392
		case 124:
			goto tr397
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st392
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st392
			}
		default:
			goto st392
		}
		goto st0
tr392:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st393
	st393:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof393
		}
	st_case_393:
//line ./combat/tokenizer.go:7644
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr395
		case 40:
			goto tr396
		case 95:
			goto st393
		case 124:
			goto tr397
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st393
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st393
			}
		default:
			goto st393
		}
		goto st0
tr27:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st394
	st394:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof394
		}
	st_case_394:
//line ./combat/tokenizer.go:7682
		if ( t.data)[( t.p)] == 97 {
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
		if ( t.data)[( t.p)] == 116 {
			goto st397
		}
		goto st0
	st397:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof397
		}
	st_case_397:
		if ( t.data)[( t.p)] == 105 {
			goto st398
		}
		goto st0
	st398:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof398
		}
	st_case_398:
		if ( t.data)[( t.p)] == 99 {
			goto st399
		}
		goto st0
	st399:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof399
		}
	st_case_399:
		if ( t.data)[( t.p)] == 105 {
			goto st400
		}
		goto st0
	st400:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof400
		}
	st_case_400:
		if ( t.data)[( t.p)] == 112 {
			goto st401
		}
		goto st0
	st401:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof401
		}
	st_case_401:
		if ( t.data)[( t.p)] == 97 {
			goto st402
		}
		goto st0
	st402:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof402
		}
	st_case_402:
		if ( t.data)[( t.p)] == 110 {
			goto st403
		}
		goto st0
	st403:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof403
		}
	st_case_403:
		if ( t.data)[( t.p)] == 116 {
			goto tr458
		}
		goto st0
tr458:
//line ./combat/tokenizer.go.rl:128
 t.tok(PARTICIPANT) 
	goto st404
	st404:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof404
		}
	st_case_404:
//line ./combat/tokenizer.go:7777
		if ( t.data)[( t.p)] == 32 {
			goto st405
		}
		goto st0
	st405:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof405
		}
	st_case_405:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st405
		case 40:
			goto tr460
		case 95:
			goto tr461
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr461
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr461
			}
		default:
			goto tr461
		}
		goto st0
tr460:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st406
	st406:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof406
		}
	st_case_406:
//line ./combat/tokenizer.go:7822
		if ( t.data)[( t.p)] == 95 {
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
	st407:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof407
		}
	st_case_407:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st408
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
	st408:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof408
		}
	st_case_408:
		if ( t.data)[( t.p)] == 9 {
			goto tr464
		}
		goto st0
tr464:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st409
	st409:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof409
		}
	st_case_409:
//line ./combat/tokenizer.go:7886
		if ( t.data)[( t.p)] == 32 {
			goto st410
		}
		goto st0
	st410:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof410
		}
	st_case_410:
		switch ( t.data)[( t.p)] {
		case 9:
			goto st534
		case 32:
			goto st459
		case 40:
			goto tr468
		case 95:
			goto tr469
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr469
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr469
			}
		default:
			goto tr469
		}
		goto st0
tr523:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st534
	st534:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof534
		}
	st_case_534:
//line ./combat/tokenizer.go:7933
		switch ( t.data)[( t.p)] {
		case 32:
			goto st535
		case 60:
			goto st411
		case 116:
			goto st421
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
	st535:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof535
		}
	st_case_535:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st522
		case 60:
			goto st411
		case 116:
			goto st421
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
	st411:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof411
		}
	st_case_411:
		switch ( t.data)[( t.p)] {
		case 70:
			goto st376
		case 98:
			goto tr470
		case 100:
			goto tr471
		case 104:
			goto tr472
		}
		goto st0
tr470:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st412
	st412:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof412
		}
	st_case_412:
//line ./combat/tokenizer.go:7993
		if ( t.data)[( t.p)] == 117 {
			goto st413
		}
		goto st0
	st413:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof413
		}
	st_case_413:
		if ( t.data)[( t.p)] == 102 {
			goto st414
		}
		goto st0
	st414:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof414
		}
	st_case_414:
		if ( t.data)[( t.p)] == 102 {
			goto st415
		}
		goto st0
	st415:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof415
		}
	st_case_415:
		if ( t.data)[( t.p)] == 62 {
			goto tr476
		}
		goto st0
tr476:
//line ./combat/tokenizer.go.rl:148
t.tokval(newAnyVal(PARTICIPATION_MODIFIER, strTok(t.data[t.prev:t.p])))
	goto st536
	st536:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof536
		}
	st_case_536:
//line ./combat/tokenizer.go:8034
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr612
		case 60:
			goto tr613
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr611
		}
		goto st0
tr612:
//line ./combat/tokenizer.go.rl:149
t.tok(PARTICIPATION_MODIFIERS_END)
	goto st537
	st537:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof537
		}
	st_case_537:
//line ./combat/tokenizer.go:8054
		switch ( t.data)[( t.p)] {
		case 32:
			goto st522
		case 60:
			goto st411
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
tr471:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st416
	st416:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof416
		}
	st_case_416:
//line ./combat/tokenizer.go:8079
		if ( t.data)[( t.p)] == 101 {
			goto st417
		}
		goto st0
	st417:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof417
		}
	st_case_417:
		if ( t.data)[( t.p)] == 98 {
			goto st412
		}
		goto st0
tr472:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st418
	st418:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof418
		}
	st_case_418:
//line ./combat/tokenizer.go:8107
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
		if ( t.data)[( t.p)] == 108 {
			goto st415
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
		if ( t.data)[( t.p)] == 116 {
			goto st423
		}
		goto st0
	st423:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof423
		}
	st_case_423:
		if ( t.data)[( t.p)] == 97 {
			goto st424
		}
		goto st0
	st424:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof424
		}
	st_case_424:
		if ( t.data)[( t.p)] == 108 {
			goto st425
		}
		goto st0
	st425:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof425
		}
	st_case_425:
		if ( t.data)[( t.p)] == 68 {
			goto st426
		}
		goto st0
	st426:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof426
		}
	st_case_426:
		if ( t.data)[( t.p)] == 97 {
			goto st427
		}
		goto st0
	st427:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof427
		}
	st_case_427:
		if ( t.data)[( t.p)] == 109 {
			goto st428
		}
		goto st0
	st428:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof428
		}
	st_case_428:
		if ( t.data)[( t.p)] == 97 {
			goto st429
		}
		goto st0
	st429:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof429
		}
	st_case_429:
		if ( t.data)[( t.p)] == 103 {
			goto st430
		}
		goto st0
	st430:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof430
		}
	st_case_430:
		if ( t.data)[( t.p)] == 101 {
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
		if ( t.data)[( t.p)] == 45 {
			goto tr492
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr493
		}
		goto st0
tr492:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st433
	st433:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof433
		}
	st_case_433:
//line ./combat/tokenizer.go:8255
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st434
		}
		goto st0
tr493:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st434
	st434:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof434
		}
	st_case_434:
//line ./combat/tokenizer.go:8274
		switch ( t.data)[( t.p)] {
		case 46:
			goto st435
		case 59:
			goto tr496
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st434
		}
		goto st0
	st435:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof435
		}
	st_case_435:
		if ( t.data)[( t.p)] == 59 {
			goto tr496
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st435
		}
		goto st0
tr496:
//line ./combat/tokenizer.go.rl:90

		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			{( t.p)++;  t.cs = 436; goto _out }
		}
		t.tokval(floatTok(float32(Float)))
	
	goto st436
	st436:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof436
		}
	st_case_436:
//line ./combat/tokenizer.go:8313
		if ( t.data)[( t.p)] == 32 {
			goto st437
		}
		goto st0
	st437:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof437
		}
	st_case_437:
		if ( t.data)[( t.p)] == 109 {
			goto st438
		}
		goto st0
	st438:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof438
		}
	st_case_438:
		if ( t.data)[( t.p)] == 111 {
			goto st439
		}
		goto st0
	st439:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof439
		}
	st_case_439:
		if ( t.data)[( t.p)] == 115 {
			goto st440
		}
		goto st0
	st440:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof440
		}
	st_case_440:
		if ( t.data)[( t.p)] == 116 {
			goto st441
		}
		goto st0
	st441:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof441
		}
	st_case_441:
		if ( t.data)[( t.p)] == 68 {
			goto st442
		}
		goto st0
	st442:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof442
		}
	st_case_442:
		if ( t.data)[( t.p)] == 97 {
			goto st443
		}
		goto st0
	st443:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof443
		}
	st_case_443:
		if ( t.data)[( t.p)] == 109 {
			goto st444
		}
		goto st0
	st444:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof444
		}
	st_case_444:
		if ( t.data)[( t.p)] == 97 {
			goto st445
		}
		goto st0
	st445:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof445
		}
	st_case_445:
		if ( t.data)[( t.p)] == 103 {
			goto st446
		}
		goto st0
	st446:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof446
		}
	st_case_446:
		if ( t.data)[( t.p)] == 101 {
			goto st447
		}
		goto st0
	st447:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof447
		}
	st_case_447:
		if ( t.data)[( t.p)] == 87 {
			goto st448
		}
		goto st0
	st448:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof448
		}
	st_case_448:
		if ( t.data)[( t.p)] == 105 {
			goto st449
		}
		goto st0
	st449:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof449
		}
	st_case_449:
		if ( t.data)[( t.p)] == 116 {
			goto st450
		}
		goto st0
	st450:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof450
		}
	st_case_450:
		if ( t.data)[( t.p)] == 104 {
			goto st451
		}
		goto st0
	st451:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof451
		}
	st_case_451:
		if ( t.data)[( t.p)] == 32 {
			goto st452
		}
		goto st0
	st452:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof452
		}
	st_case_452:
		if ( t.data)[( t.p)] == 39 {
			goto st453
		}
		goto st0
	st453:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof453
		}
	st_case_453:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr514
		case 95:
			goto tr515
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr515
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr515
			}
		default:
			goto tr515
		}
		goto st0
tr514:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st454
	st454:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof454
		}
	st_case_454:
//line ./combat/tokenizer.go:8500
		if ( t.data)[( t.p)] == 95 {
			goto st455
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st455
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	st455:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof455
		}
	st_case_455:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st456
		case 95:
			goto st455
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st455
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	st456:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof456
		}
	st_case_456:
		if ( t.data)[( t.p)] == 39 {
			goto tr518
		}
		goto st0
tr518:
//line ./combat/tokenizer.go.rl:137
t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))
	goto st457
	st457:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof457
		}
	st_case_457:
//line ./combat/tokenizer.go:8559
		if ( t.data)[( t.p)] == 59 {
			goto st538
		}
		goto st0
	st538:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof538
		}
	st_case_538:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st537
		case 60:
			goto st411
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st522
		}
		goto st0
tr515:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:8593
		switch ( t.data)[( t.p)] {
		case 39:
			goto tr518
		case 95:
			goto st458
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st458
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st458
			}
		default:
			goto st458
		}
		goto st0
tr524:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st459
	st459:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof459
		}
	st_case_459:
//line ./combat/tokenizer.go:8627
		switch ( t.data)[( t.p)] {
		case 9:
			goto st534
		case 32:
			goto st459
		}
		goto st0
tr468:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st460
	st460:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof460
		}
	st_case_460:
//line ./combat/tokenizer.go:8649
		if ( t.data)[( t.p)] == 95 {
			goto st461
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st461
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st461
			}
		default:
			goto st461
		}
		goto st0
	st461:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof461
		}
	st_case_461:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st462
		case 95:
			goto st461
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st461
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st461
			}
		default:
			goto st461
		}
		goto st0
	st462:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof462
		}
	st_case_462:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr523
		case 32:
			goto tr524
		}
		goto st0
tr469:
//line ./combat/tokenizer.go.rl:106

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
//line ./combat/tokenizer.go:8716
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr523
		case 32:
			goto tr524
		case 95:
			goto st463
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st463
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st463
			}
		default:
			goto st463
		}
		goto st0
tr461:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st464
	st464:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof464
		}
	st_case_464:
//line ./combat/tokenizer.go:8752
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr464
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
tr28:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st465
	st465:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof465
		}
	st_case_465:
//line ./combat/tokenizer.go:8786
		if ( t.data)[( t.p)] == 101 {
			goto st466
		}
		goto st0
	st466:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof466
		}
	st_case_466:
		if ( t.data)[( t.p)] == 119 {
			goto st467
		}
		goto st0
	st467:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof467
		}
	st_case_467:
		if ( t.data)[( t.p)] == 97 {
			goto st468
		}
		goto st0
	st468:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof468
		}
	st_case_468:
		if ( t.data)[( t.p)] == 114 {
			goto st469
		}
		goto st0
	st469:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof469
		}
	st_case_469:
		if ( t.data)[( t.p)] == 100 {
			goto tr531
		}
		goto st0
tr531:
//line ./combat/tokenizer.go.rl:131
 t.tok(REWARD) 
	goto st470
	st470:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof470
		}
	st_case_470:
//line ./combat/tokenizer.go:8836
		if ( t.data)[( t.p)] == 32 {
			goto st471
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st471
		}
		goto st0
	st471:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof471
		}
	st_case_471:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st471
		case 40:
			goto tr533
		case 95:
			goto tr534
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st471
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr534
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr534
			}
		default:
			goto tr534
		}
		goto st0
tr533:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st472
	st472:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof472
		}
	st_case_472:
//line ./combat/tokenizer.go:8889
		if ( t.data)[( t.p)] == 95 {
			goto st473
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st473
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st473
			}
		default:
			goto st473
		}
		goto st0
	st473:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof473
		}
	st_case_473:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st474
		case 95:
			goto st473
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st473
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st473
			}
		default:
			goto st473
		}
		goto st0
	st474:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof474
		}
	st_case_474:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr537
		case 32:
			goto tr539
		}
		if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr538
		}
		goto st0
tr537:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st475
	st475:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof475
		}
	st_case_475:
//line ./combat/tokenizer.go:8959
		switch ( t.data)[( t.p)] {
		case 32:
			goto st475
		case 45:
			goto tr541
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr542
			}
		case ( t.data)[( t.p)] >= 9:
			goto st475
		}
		goto st0
tr541:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st476
	st476:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof476
		}
	st_case_476:
//line ./combat/tokenizer.go:8989
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st477
		}
		goto st0
tr542:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st477
	st477:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof477
		}
	st_case_477:
//line ./combat/tokenizer.go:9008
		if ( t.data)[( t.p)] == 32 {
			goto tr544
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st477
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr544
		}
		goto st0
tr544:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 478; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st478
	st478:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof478
		}
	st_case_478:
//line ./combat/tokenizer.go:9037
		switch ( t.data)[( t.p)] {
		case 32:
			goto st478
		case 99:
			goto tr546
		case 101:
			goto tr547
		case 107:
			goto tr548
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st478
		}
		goto st0
tr546:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st479
	st479:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof479
		}
	st_case_479:
//line ./combat/tokenizer.go:9066
		if ( t.data)[( t.p)] == 114 {
			goto st480
		}
		goto st0
	st480:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof480
		}
	st_case_480:
		if ( t.data)[( t.p)] == 101 {
			goto st481
		}
		goto st0
	st481:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof481
		}
	st_case_481:
		if ( t.data)[( t.p)] == 100 {
			goto st482
		}
		goto st0
	st482:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof482
		}
	st_case_482:
		if ( t.data)[( t.p)] == 105 {
			goto st483
		}
		goto st0
	st483:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof483
		}
	st_case_483:
		if ( t.data)[( t.p)] == 116 {
			goto st484
		}
		goto st0
	st484:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof484
		}
	st_case_484:
		if ( t.data)[( t.p)] == 115 {
			goto st485
		}
		goto st0
	st485:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof485
		}
	st_case_485:
		if ( t.data)[( t.p)] == 32 {
			goto tr555
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr555
		}
		goto st0
tr555:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st486
	st486:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof486
		}
	st_case_486:
//line ./combat/tokenizer.go:9142
		switch ( t.data)[( t.p)] {
		case 32:
			goto st486
		case 102:
			goto tr557
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st486
		}
		goto st0
tr557:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st487
	st487:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof487
		}
	st_case_487:
//line ./combat/tokenizer.go:9167
		if ( t.data)[( t.p)] == 111 {
			goto st488
		}
		goto st0
	st488:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof488
		}
	st_case_488:
		if ( t.data)[( t.p)] == 114 {
			goto st489
		}
		goto st0
	st489:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof489
		}
	st_case_489:
		if ( t.data)[( t.p)] == 32 {
			goto st539
		}
		goto st0
	st539:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof539
		}
	st_case_539:
		switch ( t.data)[( t.p)] {
		case 10:
			goto tr616
		case 13:
			goto tr616
		}
		goto tr615
tr615:
//line ./combat/tokenizer.go.rl:171
t.tokval(strTok(t.data[t.p:]))
	goto st540
	st540:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof540
		}
	st_case_540:
//line ./combat/tokenizer.go:9211
		switch ( t.data)[( t.p)] {
		case 10:
			goto st522
		case 13:
			goto st522
		}
		goto st540
tr547:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st490
	st490:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof490
		}
	st_case_490:
//line ./combat/tokenizer.go:9233
		switch ( t.data)[( t.p)] {
		case 102:
			goto st491
		case 120:
			goto st503
		}
		goto st0
	st491:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof491
		}
	st_case_491:
		if ( t.data)[( t.p)] == 102 {
			goto st492
		}
		goto st0
	st492:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof492
		}
	st_case_492:
		if ( t.data)[( t.p)] == 101 {
			goto st493
		}
		goto st0
	st493:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof493
		}
	st_case_493:
		if ( t.data)[( t.p)] == 99 {
			goto st494
		}
		goto st0
	st494:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof494
		}
	st_case_494:
		if ( t.data)[( t.p)] == 116 {
			goto st495
		}
		goto st0
	st495:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof495
		}
	st_case_495:
		if ( t.data)[( t.p)] == 105 {
			goto st496
		}
		goto st0
	st496:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof496
		}
	st_case_496:
		if ( t.data)[( t.p)] == 118 {
			goto st497
		}
		goto st0
	st497:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof497
		}
	st_case_497:
		if ( t.data)[( t.p)] == 101 {
			goto st498
		}
		goto st0
	st498:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof498
		}
	st_case_498:
		if ( t.data)[( t.p)] == 32 {
			goto st499
		}
		goto st0
	st499:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof499
		}
	st_case_499:
		if ( t.data)[( t.p)] == 112 {
			goto st500
		}
		goto st0
	st500:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof500
		}
	st_case_500:
		if ( t.data)[( t.p)] == 111 {
			goto st501
		}
		goto st0
	st501:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof501
		}
	st_case_501:
		if ( t.data)[( t.p)] == 105 {
			goto st502
		}
		goto st0
	st502:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof502
		}
	st_case_502:
		if ( t.data)[( t.p)] == 110 {
			goto st483
		}
		goto st0
	st503:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof503
		}
	st_case_503:
		if ( t.data)[( t.p)] == 112 {
			goto st504
		}
		goto st0
	st504:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof504
		}
	st_case_504:
		if ( t.data)[( t.p)] == 101 {
			goto st505
		}
		goto st0
	st505:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof505
		}
	st_case_505:
		if ( t.data)[( t.p)] == 114 {
			goto st506
		}
		goto st0
	st506:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof506
		}
	st_case_506:
		if ( t.data)[( t.p)] == 105 {
			goto st507
		}
		goto st0
	st507:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof507
		}
	st_case_507:
		if ( t.data)[( t.p)] == 101 {
			goto st508
		}
		goto st0
	st508:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof508
		}
	st_case_508:
		if ( t.data)[( t.p)] == 110 {
			goto st509
		}
		goto st0
	st509:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof509
		}
	st_case_509:
		if ( t.data)[( t.p)] == 99 {
			goto st510
		}
		goto st0
	st510:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof510
		}
	st_case_510:
		if ( t.data)[( t.p)] == 101 {
			goto st485
		}
		goto st0
tr548:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st511
	st511:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof511
		}
	st_case_511:
//line ./combat/tokenizer.go:9435
		if ( t.data)[( t.p)] == 97 {
			goto st512
		}
		goto st0
	st512:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof512
		}
	st_case_512:
		if ( t.data)[( t.p)] == 114 {
			goto st513
		}
		goto st0
	st513:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof513
		}
	st_case_513:
		if ( t.data)[( t.p)] == 109 {
			goto st514
		}
		goto st0
	st514:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof514
		}
	st_case_514:
		if ( t.data)[( t.p)] == 97 {
			goto st485
		}
		goto st0
tr538:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st515
	st515:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof515
		}
	st_case_515:
//line ./combat/tokenizer.go:9481
		switch ( t.data)[( t.p)] {
		case 9:
			goto st475
		case 32:
			goto st515
		}
		if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st515
		}
		goto st0
tr539:
//line ./combat/tokenizer.go.rl:76

		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st516
	st516:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof516
		}
	st_case_516:
//line ./combat/tokenizer.go:9506
		switch ( t.data)[( t.p)] {
		case 9:
			goto st475
		case 32:
			goto st516
		case 40:
			goto tr586
		case 95:
			goto tr587
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto st515
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto tr587
				}
			case ( t.data)[( t.p)] >= 65:
				goto tr587
			}
		default:
			goto tr587
		}
		goto st0
tr586:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st517
	st517:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof517
		}
	st_case_517:
//line ./combat/tokenizer.go:9549
		if ( t.data)[( t.p)] == 95 {
			goto st518
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st518
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st518
			}
		default:
			goto st518
		}
		goto st0
	st518:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof518
		}
	st_case_518:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st519
		case 95:
			goto st518
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st518
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st518
			}
		default:
			goto st518
		}
		goto st0
	st519:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof519
		}
	st_case_519:
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr537
		case 32:
			goto tr538
		}
		if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr538
		}
		goto st0
tr587:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st520
	st520:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof520
		}
	st_case_520:
//line ./combat/tokenizer.go:9619
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr537
		case 32:
			goto tr538
		case 95:
			goto st520
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr538
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st520
				}
			case ( t.data)[( t.p)] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st0
tr534:
//line ./combat/tokenizer.go.rl:106

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st521
	st521:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof521
		}
	st_case_521:
//line ./combat/tokenizer.go:9660
		switch ( t.data)[( t.p)] {
		case 9:
			goto tr537
		case 32:
			goto tr539
		case 95:
			goto st521
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 10 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr538
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st521
				}
			case ( t.data)[( t.p)] >= 65:
				goto st521
			}
		default:
			goto st521
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
	_test_eof522:  t.cs = 522; goto _test_eof
	_test_eof63:  t.cs = 63; goto _test_eof
	_test_eof64:  t.cs = 64; goto _test_eof
	_test_eof65:  t.cs = 65; goto _test_eof
	_test_eof66:  t.cs = 66; goto _test_eof
	_test_eof67:  t.cs = 67; goto _test_eof
	_test_eof68:  t.cs = 68; goto _test_eof
	_test_eof69:  t.cs = 69; goto _test_eof
	_test_eof70:  t.cs = 70; goto _test_eof
	_test_eof71:  t.cs = 71; goto _test_eof
	_test_eof72:  t.cs = 72; goto _test_eof
	_test_eof73:  t.cs = 73; goto _test_eof
	_test_eof74:  t.cs = 74; goto _test_eof
	_test_eof75:  t.cs = 75; goto _test_eof
	_test_eof76:  t.cs = 76; goto _test_eof
	_test_eof77:  t.cs = 77; goto _test_eof
	_test_eof78:  t.cs = 78; goto _test_eof
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
	_test_eof523:  t.cs = 523; goto _test_eof
	_test_eof524:  t.cs = 524; goto _test_eof
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
	_test_eof525:  t.cs = 525; goto _test_eof
	_test_eof526:  t.cs = 526; goto _test_eof
	_test_eof191:  t.cs = 191; goto _test_eof
	_test_eof192:  t.cs = 192; goto _test_eof
	_test_eof193:  t.cs = 193; goto _test_eof
	_test_eof194:  t.cs = 194; goto _test_eof
	_test_eof195:  t.cs = 195; goto _test_eof
	_test_eof196:  t.cs = 196; goto _test_eof
	_test_eof197:  t.cs = 197; goto _test_eof
	_test_eof198:  t.cs = 198; goto _test_eof
	_test_eof527:  t.cs = 527; goto _test_eof
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
	_test_eof528:  t.cs = 528; goto _test_eof
	_test_eof529:  t.cs = 529; goto _test_eof
	_test_eof333:  t.cs = 333; goto _test_eof
	_test_eof334:  t.cs = 334; goto _test_eof
	_test_eof335:  t.cs = 335; goto _test_eof
	_test_eof336:  t.cs = 336; goto _test_eof
	_test_eof337:  t.cs = 337; goto _test_eof
	_test_eof338:  t.cs = 338; goto _test_eof
	_test_eof339:  t.cs = 339; goto _test_eof
	_test_eof340:  t.cs = 340; goto _test_eof
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
	_test_eof530:  t.cs = 530; goto _test_eof
	_test_eof531:  t.cs = 531; goto _test_eof
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
	_test_eof532:  t.cs = 532; goto _test_eof
	_test_eof533:  t.cs = 533; goto _test_eof
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
	_test_eof534:  t.cs = 534; goto _test_eof
	_test_eof535:  t.cs = 535; goto _test_eof
	_test_eof411:  t.cs = 411; goto _test_eof
	_test_eof412:  t.cs = 412; goto _test_eof
	_test_eof413:  t.cs = 413; goto _test_eof
	_test_eof414:  t.cs = 414; goto _test_eof
	_test_eof415:  t.cs = 415; goto _test_eof
	_test_eof536:  t.cs = 536; goto _test_eof
	_test_eof537:  t.cs = 537; goto _test_eof
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
	_test_eof538:  t.cs = 538; goto _test_eof
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
	_test_eof539:  t.cs = 539; goto _test_eof
	_test_eof540:  t.cs = 540; goto _test_eof
	_test_eof490:  t.cs = 490; goto _test_eof
	_test_eof491:  t.cs = 491; goto _test_eof
	_test_eof492:  t.cs = 492; goto _test_eof
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

	_test_eof: {}
	if ( t.p) == ( t.pe) {
		switch  t.cs {
		case 522, 524, 525, 526, 530, 531, 534, 535, 537, 538, 540:
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
		case 527:
//line ./combat/tokenizer.go.rl:82

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 0; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
		case 528, 529, 532, 533:
//line ./combat/tokenizer.go.rl:137
t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
		case 523:
//line ./combat/tokenizer.go.rl:138
t.tokval(newAnyVal(DAMAGE_MODIFIER, strTok(t.data[t.prev:t.p])))
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
		case 536:
//line ./combat/tokenizer.go.rl:149
t.tok(PARTICIPATION_MODIFIERS_END)
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
		case 539:
//line ./combat/tokenizer.go.rl:171
t.tokval(strTok(t.data[t.p:]))
//line ./combat/tokenizer.go.rl:119
 t.tok(EOL) 
//line ./combat/tokenizer.go:10266
		}
	}

	_out: {}
	}

//line ./combat/tokenizer.go.rl:203

	if debugTokenizer {
		fmt.Printf("exited: %s\n", t.state)
	}
	if t.p != t.pe || (len(t.tokens) > 0 && t.tokens[len(t.tokens)-1] != voidTok(EOL)) {
		t.err(fmt.Errorf("%w: %q %q", ErrLineIsNotFinished, t.data[:t.p], t.data[t.p:t.pe]))
	}

	return t.tokens, errors.Join(t.errors...)
}

func (t *tokenizer) tokval(token token) {
	t.tokens = append(t.tokens, token)
}

func (t *tokenizer) tok(tok int) {
	t.tokens = append(t.tokens, voidTok(tok))
}

func (t *tokenizer) err(err error) {
	t.errors = append(t.errors, err)
}