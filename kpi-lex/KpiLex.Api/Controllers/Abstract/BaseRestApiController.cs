using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.BusinessLogic.Services.Abstract;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers.Abstract
{
    public abstract class BaseRestApiController<TRequestModel, TResponseModel, TDomainModel, TIdentity> : Controller
    {
        protected readonly IService<TDomainModel, TIdentity> Service;
        protected readonly IApiMapper<TDomainModel, TResponseModel> ResponseModelMapper;
        protected readonly IApiMapper<TRequestModel, TDomainModel> RequestModelMapper;

        protected BaseRestApiController(
            IService<TDomainModel, TIdentity> service,
            IApiMapper<TDomainModel, TResponseModel> responseModelMapper,
            IApiMapper<TRequestModel, TDomainModel> requestModelMapper)
        {
            Service = service;
            ResponseModelMapper = responseModelMapper;
            RequestModelMapper = requestModelMapper;
        }
        
        [HttpGet]
        public virtual async Task<IReadOnlyCollection<TResponseModel>> Get()
        {
            var entities = await Service.GetAll();
            return entities.Select(c => ResponseModelMapper.Map(c)).ToList();
        }

        [HttpGet, Route("{id}")]
        public virtual async Task<TResponseModel> Get([FromRoute] TIdentity id)
        {
            var entity = await Service.GetById(id);
            return ResponseModelMapper.Map(entity);
        } 
        
        [HttpPost]
        public virtual async Task<TResponseModel> Add(
            [FromBody] TRequestModel requestModel)
        {
            var model = RequestModelMapper.Map(requestModel);
            model = await Service.AddAsync(model);
            return ResponseModelMapper.Map(model);
        }
        
        [HttpPut]
        public virtual async Task<IActionResult> Update(
            [FromBody] TRequestModel requestModel)
        {
            var model = RequestModelMapper.Map(requestModel);
            await Service.UpdateAsync(model);
            return Ok();
        }
        
        [HttpDelete, Route("{id}")]
        public virtual async Task<IActionResult> Update(
            [FromBody] TIdentity id)
        {
            await Service.RemoveAsync(id);
            return Ok();
        }
    }
}