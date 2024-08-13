package jp.ac.tohoku.qse.takahashi.AtcSimulator.domain.model.aggregate.airspace;

import jp.ac.tohoku.qse.takahashi.AtcSimulator.domain.model.entity.aircraft.Aircraft;

public interface AirSpace {
    public void addAircraft(Aircraft aircraft);
    public void removeAircraft(Aircraft aircraft);

    void NextStep();
}