using KpiLex.Api.Constants;
using KpiLex.Api.Controllers.Abstract;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Routing;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.FacultyRoute)]
    public class FacultyController : BaseRestApiController<FacultyRequestModel, FacultyResponseModel, Faculty, int>
    {
        public FacultyController(IService<Faculty, int> service,
            IApiMapper<Faculty, FacultyResponseModel> responseModelMapper,
            IApiMapper<FacultyRequestModel, Faculty> requestModelMapper)
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}