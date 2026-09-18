package smartfolder

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niossmartfolder "github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/smartfolder"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/smartfolder"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type SmartfolderPersonalService interface {
	Create(ctx context.Context, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error)
	Update(ctx context.Context, id string, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*smartfolder.SmartfolderPersonal, *http.Response, string, error)
}

type smartfolderPersonalService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSmartfolderPersonalService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SmartfolderPersonalService {
	return &smartfolderPersonalService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SmartfolderPersonal and returns the created object
func (s *smartfolderPersonalService) Create(ctx context.Context, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *smartfolderPersonalService) createNIOS(ctx context.Context, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	payload, err := common.MapTo[niossmartfolder.SmartfolderPersonal](obj, mapper.SmartfolderPersonalNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SmartFolderAPI.SmartfolderPersonalAPI.
		Create(ctx).
		SmartfolderPersonal(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSmartfolderPersonalResponseAsObject.GetResult()

	return mapNIOSSmartfolderPersonalToResponse(&result), httpResp, nil
}

// Read retrieves a SmartfolderPersonal by ID
func (s *smartfolderPersonalService) Read(ctx context.Context, id string, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *smartfolderPersonalService) readNIOS(ctx context.Context, id string, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	req := s.niosClient.SmartFolderAPI.SmartfolderPersonalAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSmartfolderPersonalResponseObjectAsResult.GetResult()

	return mapNIOSSmartfolderPersonalToResponse(&result), httpResp, nil
}

// Update modifies an existing SmartfolderPersonal and returns the updated object
func (s *smartfolderPersonalService) Update(ctx context.Context, id string, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *smartfolderPersonalService) updateNIOS(ctx context.Context, id string, obj *smartfolder.SmartfolderPersonal, opts *core.Options) (*smartfolder.SmartfolderPersonal, *http.Response, error) {
	payload, err := common.MapTo[niossmartfolder.SmartfolderPersonal](obj, mapper.SmartfolderPersonalNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.SmartFolderAPI.SmartfolderPersonalAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SmartfolderPersonal(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSmartfolderPersonalResponseAsObject.GetResult()

	return mapNIOSSmartfolderPersonalToResponse(&result), httpResp, nil
}

// Delete removes a SmartfolderPersonal by ID
func (s *smartfolderPersonalService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *smartfolderPersonalService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.SmartFolderAPI.SmartfolderPersonalAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SmartfolderPersonal objects based on filter options
func (s *smartfolderPersonalService) List(ctx context.Context, opts *core.ListOptions) ([]*smartfolder.SmartfolderPersonal, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *smartfolderPersonalService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*smartfolder.SmartfolderPersonal, *http.Response, string, error) {
	req := s.niosClient.SmartFolderAPI.SmartfolderPersonalAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SmartfolderPersonalFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSmartfolderPersonalResponseObject.GetResult()
	items := make([]*smartfolder.SmartfolderPersonal, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSmartfolderPersonalToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSmartfolderPersonalResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSmartfolderPersonalToResponse(r *niossmartfolder.SmartfolderPersonal) *smartfolder.SmartfolderPersonal {
	resp := &smartfolder.SmartfolderPersonal{
		Id: r.Ref,
	}
	resp.NIOS = &smartfolder.NIOSSmartfolderPersonalExt{
		Comment:    r.Comment,
		GroupBys:   r.GroupBys,
		Name:       r.Name,
		QueryItems: r.QueryItems,
	}
	return resp
}
