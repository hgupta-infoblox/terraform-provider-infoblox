package security

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RadiusAuthserviceService interface {
	Create(ctx context.Context, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error)
	Update(ctx context.Context, id string, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*security.RadiusAuthservice, *http.Response, string, error)
}

type radiusAuthserviceService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRadiusAuthserviceService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RadiusAuthserviceService {
	return &radiusAuthserviceService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RadiusAuthservice and returns the created object
func (s *radiusAuthserviceService) Create(ctx context.Context, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *radiusAuthserviceService) createNIOS(ctx context.Context, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.RadiusAuthservice](obj, mapper.RadiusAuthserviceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SecurityAPI.RadiusAuthserviceAPI.
		Create(ctx).
		RadiusAuthservice(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRadiusAuthserviceResponseAsObject.GetResult()

	return mapNIOSRadiusAuthserviceToResponse(&result), httpResp, nil
}

// Read retrieves a RadiusAuthservice by ID
func (s *radiusAuthserviceService) Read(ctx context.Context, id string, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *radiusAuthserviceService) readNIOS(ctx context.Context, id string, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	req := s.niosClient.SecurityAPI.RadiusAuthserviceAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRadiusAuthserviceResponseObjectAsResult.GetResult()

	return mapNIOSRadiusAuthserviceToResponse(&result), httpResp, nil
}

// Update modifies an existing RadiusAuthservice and returns the updated object
func (s *radiusAuthserviceService) Update(ctx context.Context, id string, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *radiusAuthserviceService) updateNIOS(ctx context.Context, id string, obj *security.RadiusAuthservice, opts *core.Options) (*security.RadiusAuthservice, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.RadiusAuthservice](obj, mapper.RadiusAuthserviceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SecurityAPI.RadiusAuthserviceAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RadiusAuthservice(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRadiusAuthserviceResponseAsObject.GetResult()

	return mapNIOSRadiusAuthserviceToResponse(&result), httpResp, nil
}

// Delete removes a RadiusAuthservice by ID
func (s *radiusAuthserviceService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *radiusAuthserviceService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.SecurityAPI.RadiusAuthserviceAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RadiusAuthservice objects based on filter options
func (s *radiusAuthserviceService) List(ctx context.Context, opts *core.ListOptions) ([]*security.RadiusAuthservice, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *radiusAuthserviceService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*security.RadiusAuthservice, *http.Response, string, error) {
	req := s.niosClient.SecurityAPI.RadiusAuthserviceAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RadiusAuthserviceFilterFieldMap[core.BackendNIOS])
			filters := make(map[string]any, len(translatedFilters))
			for k, v := range translatedFilters {
				filters[k] = v
			}
			req = req.Filters(filters)
		}
		if len(opts.ExtAttrFilter) > 0 {
			extAttrFilters := make(map[string]any, len(opts.ExtAttrFilter))
			for k, v := range opts.ExtAttrFilter {
				extAttrFilters[k] = v
			}
			req = req.Extattrfilter(extAttrFilters)
		}
		if opts.PageID != "" {
			req = req.PageId(opts.PageID)
		}
		req = req.Paging(opts.Paging)
		maxResults := opts.MaxResults
		if maxResults <= 0 {
			maxResults = core.DefaultListLimit
		}
		req = req.MaxResults(maxResults)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListRadiusAuthserviceResponseObject.GetResult()
	items := make([]*security.RadiusAuthservice, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRadiusAuthserviceToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRadiusAuthserviceResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRadiusAuthserviceToResponse(r *niossecurity.RadiusAuthservice) *security.RadiusAuthservice {
	resp := &security.RadiusAuthservice{
		Id: r.Ref,
	}
	resp.NIOS = &security.NIOSRadiusAuthserviceExt{
		AcctRetries:      r.AcctRetries,
		AcctTimeout:      r.AcctTimeout,
		AuthRetries:      r.AuthRetries,
		AuthTimeout:      r.AuthTimeout,
		CacheTtl:         r.CacheTtl,
		Comment:          r.Comment,
		Disable:          r.Disable,
		EnableCache:      r.EnableCache,
		Mode:             r.Mode,
		Name:             r.Name,
		RecoveryInterval: r.RecoveryInterval,
		Servers:          r.Servers,
	}
	return resp
}
