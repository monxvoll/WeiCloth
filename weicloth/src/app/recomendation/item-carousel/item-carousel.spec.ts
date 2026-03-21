import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ItemCarousel } from './item-carousel';

describe('ItemCarousel', () => {
  let component: ItemCarousel;
  let fixture: ComponentFixture<ItemCarousel>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ItemCarousel],
    }).compileComponents();

    fixture = TestBed.createComponent(ItemCarousel);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
