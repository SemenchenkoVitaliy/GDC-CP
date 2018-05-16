using KpiLex.Api.Constants;
using KpiLex.Api.Controllers.Abstract;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.SpecialitiesRoute)]
    public class
        SpecialityController : BaseRestApiController<SpecialityRequestModel, SpecialityResponseModel, Speciality, int>
    {
        public SpecialityController(IService<Speciality, int> service,
            IApiMapper<Speciality, SpecialityResponseModel> responseModelMapper,
            IApiMapper<SpecialityRequestModel, Speciality> requestModelMapper)
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}