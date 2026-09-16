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

type TacacsplusAuthserviceService interface {
	Create(ctx context.Context, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error)
	Update(ctx context.Context, id string, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*security.TacacsplusAuthservice, *http.Response, string, error)
}

type tacacsplusAuthserviceService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewTacacsplusAuthserviceService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) TacacsplusAuthserviceService {
	return &tacacsplusAuthserviceService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new TacacsplusAuthservice and returns the created object
func (s *tacacsplusAuthserviceService) Create(ctx context.Context, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tacacsplusAuthserviceService) createNIOS(ctx context.Context, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.TacacsplusAuthservice](obj, mapper.TacacsplusAuthserviceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SecurityAPI.TacacsplusAuthserviceAPI.
		Create(ctx).
		TacacsplusAuthservice(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateTacacsplusAuthserviceResponseAsObject.GetResult()

	return mapNIOSTacacsplusAuthserviceToResponse(&result), httpResp, nil
}

// Read retrieves a TacacsplusAuthservice by ID
func (s *tacacsplusAuthserviceService) Read(ctx context.Context, id string, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tacacsplusAuthserviceService) readNIOS(ctx context.Context, id string, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	req := s.niosClient.SecurityAPI.TacacsplusAuthserviceAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetTacacsplusAuthserviceResponseObjectAsResult.GetResult()

	return mapNIOSTacacsplusAuthserviceToResponse(&result), httpResp, nil
}

// Update modifies an existing TacacsplusAuthservice and returns the updated object
func (s *tacacsplusAuthserviceService) Update(ctx context.Context, id string, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tacacsplusAuthserviceService) updateNIOS(ctx context.Context, id string, obj *security.TacacsplusAuthservice, opts *core.Options) (*security.TacacsplusAuthservice, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.TacacsplusAuthservice](obj, mapper.TacacsplusAuthserviceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SecurityAPI.TacacsplusAuthserviceAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		TacacsplusAuthservice(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateTacacsplusAuthserviceResponseAsObject.GetResult()

	return mapNIOSTacacsplusAuthserviceToResponse(&result), httpResp, nil
}

// Delete removes a TacacsplusAuthservice by ID
func (s *tacacsplusAuthserviceService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tacacsplusAuthserviceService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.SecurityAPI.TacacsplusAuthserviceAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves TacacsplusAuthservice objects based on filter options
func (s *tacacsplusAuthserviceService) List(ctx context.Context, opts *core.ListOptions) ([]*security.TacacsplusAuthservice, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tacacsplusAuthserviceService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*security.TacacsplusAuthservice, *http.Response, string, error) {
	req := s.niosClient.SecurityAPI.TacacsplusAuthserviceAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.TacacsplusAuthserviceFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListTacacsplusAuthserviceResponseObject.GetResult()
	items := make([]*security.TacacsplusAuthservice, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSTacacsplusAuthserviceToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListTacacsplusAuthserviceResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSTacacsplusAuthserviceToResponse(r *niossecurity.TacacsplusAuthservice) *security.TacacsplusAuthservice {
	resp := &security.TacacsplusAuthservice{
		Id: r.Ref,
	}
	resp.NIOS = &security.NIOSTacacsplusAuthserviceExt{
		AcctRetries: r.AcctRetries,
		AcctTimeout: r.AcctTimeout,
		AuthRetries: r.AuthRetries,
		AuthTimeout: r.AuthTimeout,
		Comment:     r.Comment,
		Disable:     r.Disable,
		Name:        r.Name,
		Servers:     r.Servers,
	}
	return resp
}
