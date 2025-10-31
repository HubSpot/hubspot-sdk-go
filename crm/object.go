// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ObjectService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options             []option.RequestOption
	Calls               ObjectCallService
	Carts               ObjectCartService
	CommercePayments    ObjectCommercePaymentService
	Communications      ObjectCommunicationService
	Companies           ObjectCompanyService
	Contacts            ObjectContactService
	Contracts           ObjectContractService
	Courses             ObjectCourseService
	Custom              ObjectCustomService
	DealSplits          ObjectDealSplitService
	Deals               ObjectDealService
	Discounts           ObjectDiscountService
	Emails              ObjectEmailService
	FeedbackSubmissions ObjectFeedbackSubmissionService
	Fees                ObjectFeeService
	GoalTargets         ObjectGoalTargetService
	Invoices            ObjectInvoiceService
	Leads               ObjectLeadService
	LineItems           ObjectLineItemService
	Listings            ObjectListingService
	Meetings            ObjectMeetingService
	Notes               ObjectNoteService
	Objects             ObjectObjectService
	Orders              ObjectOrderService
	PartnerClients      ObjectPartnerClientService
	PartnerServices     ObjectPartnerServiceService
	PostalMail          ObjectPostalMailService
	Products            ObjectProductService
	Quotes              ObjectQuoteService
	Schemas             ObjectSchemaService
	Services            ObjectServiceService
	Tasks               ObjectTaskService
	Taxes               ObjectTaxService
	Tickets             ObjectTicketService
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.Options = opts
	r.Calls = NewObjectCallService(opts...)
	r.Carts = NewObjectCartService(opts...)
	r.CommercePayments = NewObjectCommercePaymentService(opts...)
	r.Communications = NewObjectCommunicationService(opts...)
	r.Companies = NewObjectCompanyService(opts...)
	r.Contacts = NewObjectContactService(opts...)
	r.Contracts = NewObjectContractService(opts...)
	r.Courses = NewObjectCourseService(opts...)
	r.Custom = NewObjectCustomService(opts...)
	r.DealSplits = NewObjectDealSplitService(opts...)
	r.Deals = NewObjectDealService(opts...)
	r.Discounts = NewObjectDiscountService(opts...)
	r.Emails = NewObjectEmailService(opts...)
	r.FeedbackSubmissions = NewObjectFeedbackSubmissionService(opts...)
	r.Fees = NewObjectFeeService(opts...)
	r.GoalTargets = NewObjectGoalTargetService(opts...)
	r.Invoices = NewObjectInvoiceService(opts...)
	r.Leads = NewObjectLeadService(opts...)
	r.LineItems = NewObjectLineItemService(opts...)
	r.Listings = NewObjectListingService(opts...)
	r.Meetings = NewObjectMeetingService(opts...)
	r.Notes = NewObjectNoteService(opts...)
	r.Objects = NewObjectObjectService(opts...)
	r.Orders = NewObjectOrderService(opts...)
	r.PartnerClients = NewObjectPartnerClientService(opts...)
	r.PartnerServices = NewObjectPartnerServiceService(opts...)
	r.PostalMail = NewObjectPostalMailService(opts...)
	r.Products = NewObjectProductService(opts...)
	r.Quotes = NewObjectQuoteService(opts...)
	r.Schemas = NewObjectSchemaService(opts...)
	r.Services = NewObjectServiceService(opts...)
	r.Tasks = NewObjectTaskService(opts...)
	r.Taxes = NewObjectTaxService(opts...)
	r.Tickets = NewObjectTicketService(opts...)
	return
}
