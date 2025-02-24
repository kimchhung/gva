package relay

import "backend/internal/relay/relayt"

func validationAndDecode(opt *PaginateGlobalConfig) (err error) {
	if opt.First != nil && opt.Last != nil {
		return relayt.NewErrRelay("passing both `first` and `last` to paginate a connection is not supported")
	}

	if opt.First != nil && *opt.First < 0 {
		return relayt.NewErrRelay("`first` on a connection cannot be less than zero")
	}

	if opt.Last != nil && *opt.Last < 0 {
		return relayt.NewErrRelay("`last` on a connection cannot be less than zero")
	}

	// if opt.After != nil {
	// 	if opt.afterRaw, err = base64.StdEncoding.DecodeString(*opt.After); err != nil {
	// 		return relayt.NewErrRelay(fmt.Sprintf("invalid `after` cursor: `%s`", *opt.After))
	// 	}
	// }

	// if opt.Before != nil {
	// 	if opt.beforeRaw, err = base64.StdEncoding.DecodeString(*opt.Before); err != nil {
	// 		return relayt.NewErrRelay(fmt.Sprintf("invalid `before` cursor: `%s`", *opt.Before))
	// 	}
	// }

	return nil
}
