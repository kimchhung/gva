package relay

func (p *PageInfo) SetHasPreviousPage(totalCount, edgesLen int, option *PaginateGlobalConfig) {
	if !option.isDisableCount && totalCount == 0 {
		return
	}

	if totalCount == edgesLen {
		return
	}

	if option.After != nil {
		p.HasPreviousPage = true
		return
	}
}

func (p *PageInfo) SetHasNextPage(totalCount, edgesLen int, option *PaginateGlobalConfig) {

	isEnableCount := !option.isDisableCount
	if isEnableCount {
		if totalCount == 0 {
			return
		}

		if totalCount == edgesLen {
			return
		}
	}

	if option.Before != nil {
		p.HasNextPage = true
		return
	}

	if option.First != nil && *option.First == edgesLen {
		p.HasNextPage = true
		return
	}

	if option.Last != nil && *option.Last == edgesLen {
		p.HasNextPage = true
		return
	}

	if option.First != nil && *option.First > edgesLen {
		return
	}

	if option.Last != nil && *option.Last > edgesLen {
		return
	}

	if isEnableCount {
		if option.After == nil && option.First != nil && *option.First < totalCount {
			p.HasNextPage = true
			return
		}

		if option.After == nil && option.Last != nil && *option.Last < totalCount {
			p.HasNextPage = true
			return
		}
	}

}
