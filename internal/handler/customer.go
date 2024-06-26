package handler

import (
	"net/http"

	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_validator"
)

func (h *Handler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterCustomer
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.RegisterCustomer(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.OK(w, http.StatusCreated, res)
}

func (h *Handler) GetListCustomer(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListCustomerQuery{
		Phone: StringPtr(qp.Get("phoneNumber")),
		Name:  StringPtr(qp.Get("name")),
	}

	res, err := h.uc.GetListCustomer(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusOK, res)

}
