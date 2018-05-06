using System;
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
    [Route(RouteConstants.LectureRoute)]
    public class LectureController : BaseRestApiController<LectureRequestModel, LectureResponseModel, Lecture, Guid>
    {
        public LectureController(IService<Lecture, Guid> service,
            IApiMapper<Lecture, LectureResponseModel> responseModelMapper,
            IApiMapper<LectureRequestModel, Lecture> requestModelMapper) 
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}