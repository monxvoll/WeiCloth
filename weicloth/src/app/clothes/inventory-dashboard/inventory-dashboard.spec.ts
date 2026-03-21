import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InventoryDashboard } from './inventory-dashboard';

describe('InventoryDashboard', () => {
  let component: InventoryDashboard;
  let fixture: ComponentFixture<InventoryDashboard>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InventoryDashboard],
    }).compileComponents();

    fixture = TestBed.createComponent(InventoryDashboard);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
