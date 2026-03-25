import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ItemCarousel } from './item-carousel/item-carousel';

export const routes: Routes = [
  { path : 'carousel', component: ItemCarousel}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule] 
})
export class RecomendationRoutingModule {}