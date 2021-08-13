package redis

// var (
// 	client *redis.Client
// )

// var (
// 	key = "key"
// 	val = "val"
// )

// func TestMain(m *testing.M) {
// 	mr, err := miniredis.Run()
// 	if err != nil {
// 		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	client = redis.NewClient(&redis.Options{
// 		Addr: mr.Addr(),
// 	})

// 	code := m.Run()
// 	os.Exit(code)
// }

// func TestConnect(t *testing.T) {
// 	Storage.setClient(client)

// 	tests := []struct {
// 		step string
// 		conf conf.AppConfig
// 		err  error
// 	}{
// 		{
// 			step: "A",
// 			conf: conf.AppConfig{
// 				Debug: true,
// 			},
// 			err: fmt.Errorf("dial tcp 127.0.0.1:6379: connect: connection refused"),
// 		}}

// 	for _, tc := range tests {
// 		t.Run(tc.step, func(t *testing.T) {
// 			once = sync.Once{}
// 			err := Storage.Connect(tc.conf)
// 			if tc.err != nil {
// 				assert.Equal(t, tc.err.Error(), err.Error())
// 			}
// 		})
// 	}
// }

// func TestSetClient(t *testing.T) {
// 	Storage.setClient(client)
// }

// func TestSet(t *testing.T) {
// 	Storage.setClient(client)

// 	if err := Storage.Set(context.Background(), key, val, time.Hour); err != nil {
// 		assert.Equal(t, nil, err)
// 	}

// 	ch := make(chan string)
// 	if err := Storage.Set(context.Background(), key, ch, time.Hour); err != nil {
// 		assert.Equal(t, "json: unsupported type: chan string", err.Error())
// 	}
// }

// func TestGet(t *testing.T) {
// 	tests := []struct {
// 		step string
// 		key  string
// 		val  string
// 		err  error
// 	}{
// 		{
// 			step: "A",
// 			err:  nil,
// 			key:  key,
// 			val:  val,
// 		},
// 		{
// 			step: "B",
// 			err:  fmt.Errorf("redis: nil"),
// 			key:  "invalid",
// 			val:  "",
// 		},
// 	}

// 	if err := Storage.Set(context.Background(), key, val, time.Hour); err != nil {
// 		assert.Equal(t, nil, err)
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.step, func(t *testing.T) {
// 			var des string
// 			if err := Storage.Get(context.Background(), tc.key, &des); err != nil {
// 				assert.Equal(t, tc.err.Error(), err.Error())
// 			}

// 			if des != tc.val {
// 				assert.Equal(t, tc.val, des)
// 			}

// 		})
// 	}
// }

// func TestDel(t *testing.T) {
// 	Storage.setClient(client)

// 	ctx, c := context.WithCancel(context.Background())
// 	c()

// 	if err := Storage.Del(ctx, key); err != nil {
// 		assert.Equal(t, "context canceled", err.Error())
// 	}

// 	if err := Storage.Del(context.Background(), key); err != nil {
// 		assert.Equal(t, nil, err)
// 	}

// }
